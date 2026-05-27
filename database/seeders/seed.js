const { MongoClient } = require('mongodb');
const fs = require('fs');
const path = require('path');
const bcrypt = require('bcryptjs');

const uri = process.env.MONGODB_URI || 'mongodb://localhost:27017';
const dbName = process.env.MONGODB_DB_NAME || 'anime_streaming_db';

async function seedDatabase() {
  const client = new MongoClient(uri);

  try {
    await client.connect();
    console.log('✅ Connected to MongoDB');

    const db = client.db(dbName);

    // Clear existing data
    await db.collection('animes').deleteMany({});
    await db.collection('episodes').deleteMany({});
    await db.collection('users').deleteMany({});
    await db.collection('comments').deleteMany({});
    await db.collection('history').deleteMany({});

    console.log('🗑️  Cleared existing collections');

    // Seed animes
    const animesData = JSON.parse(fs.readFileSync(path.join(__dirname, 'animes.json'), 'utf8'));
    const animesWithDates = animesData.map(anime => ({
      ...anime,
      episodesCount: 0,
      createdAt: new Date(),
      updatedAt: new Date()
    }));
    const animesResult = await db.collection('animes').insertMany(animesWithDates);
    console.log(`✅ Inserted ${animesResult.insertedCount} animes`);

    // Create anime slug to ID mapping
    const animeMap = {};
    const animesCursor = await db.collection('animes').find({}).toArray();
    animesCursor.forEach(anime => {
      animeMap[anime.slug] = anime._id;
    });

    // Seed episodes
    const episodesData = JSON.parse(fs.readFileSync(path.join(__dirname, 'episodes.json'), 'utf8'));
    const episodesWithRefs = episodesData.map(ep => ({
      ...ep,
      animeId: animeMap[ep.animeSlug],
      createdAt: new Date(),
      updatedAt: new Date()
    })).map(({ animeSlug, ...rest }) => rest);

    const episodesResult = await db.collection('episodes').insertMany(episodesWithRefs);
    console.log(`✅ Inserted ${episodesResult.insertedCount} episodes`);

    // Update episode counts
    for (const [slug, id] of Object.entries(animeMap)) {
      const count = await db.collection('episodes').countDocuments({ animeId: id });
      await db.collection('animes').updateOne(
        { _id: id },
        { $set: { episodesCount: count } }
      );
    }
    console.log('✅ Updated episode counts');

    // Seed users
    const usersData = JSON.parse(fs.readFileSync(path.join(__dirname, 'users.json'), 'utf8'));
    const usersWithHashes = await Promise.all(
      usersData.map(async user => ({
        ...user,
        passwordHash: await bcrypt.hash(user.password, 10),
        isActive: true,
        createdAt: new Date(),
        updatedAt: new Date()
      }))
    );
    const usersResult = await db.collection('users').insertMany(
      usersWithHashes.map(({ password, ...rest }) => rest)
    );
    console.log(`✅ Inserted ${usersResult.insertedCount} users`);

    // Create username to ID mapping
    const userMap = {};
    const usersCursor = await db.collection('users').find({}).toArray();
    usersCursor.forEach(user => {
      userMap[user.username] = user._id;
    });

    // Seed comments
    const commentsData = JSON.parse(fs.readFileSync(path.join(__dirname, 'comments.json'), 'utf8'));
    const commentsWithRefs = commentsData.map(comment => ({
      animeId: animeMap[comment.animeSlug],
      userId: userMap[comment.username],
      content: comment.content,
      likes: comment.likes,
      likedBy: [],
      isDeleted: false,
      createdAt: new Date(),
      updatedAt: new Date()
    }));

    const commentsResult = await db.collection('comments').insertMany(commentsWithRefs);
    console.log(`✅ Inserted ${commentsResult.insertedCount} comments`);

    console.log('\n🎉 Database seeded successfully!');
    console.log('\nDefault credentials:');
    console.log('  Admin: admin@animestream.com / admin123');
    console.log('  User:  user@animestream.com / user123');

  } catch (error) {
    console.error('❌ Error seeding database:', error);
  } finally {
    await client.close();
  }
}

seedDatabase();
