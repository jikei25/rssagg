package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jikei25/rssagg/internal/database"
)

type User struct {
	ID 			uuid.UUID 	`json:"id"`
	Created_at 	time.Time 	`json:"created_at"`
	Updated_at 	time.Time 	`json:"updated_at"`
	Name 		string 		`json:"name"`
	APIKey 		string 		`json:"api_key"`
}

type Feed struct {
	ID 			uuid.UUID 	`json:"id"`
	Created_at 	time.Time 	`json:"created_at"`
	Updated_at 	time.Time 	`json:"updated_at"`
	Name 		string 		`json:"name"`
	Url 		string 		`json:"url"`
	UserID 		uuid.UUID 	`json:"user_id"`
}

type FeedFollow struct {
	ID 			uuid.UUID 	`json:"id"`
	Created_at 	time.Time 	`json:"created_at"`
	Updated_at 	time.Time 	`json:"updated_at"`
	UserID 		uuid.UUID 	`json:"user_id"`
	FeedID 		uuid.UUID 	`json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID 		`json:"id"`
	CreatedAt   time.Time 		`json:"created_at"`
	UpdatedAt   time.Time 		`json:"updated_at"`
	Title       string			`json:"title"`	
	Description *string 		`json:"description"`
	PublishedAt time.Time		`json:"published_at"`
	Url         string			`json:"url"`
	FeedID      uuid.UUID		`json:"feed_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User {
		ID: dbUser.ID,
		Created_at: dbUser.CreatedAt,
		Updated_at: dbUser.UpdatedAt,
		Name: dbUser.Name,
		APIKey: dbUser.ApiKey,
	}
}


func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		Created_at: dbFeed.CreatedAt,
		Updated_at: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: dbFeedFollow.ID,
		Created_at: dbFeedFollow.CreatedAt,
		Updated_at: dbFeedFollow.UpdatedAt,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feed_follows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feed_follows = append(feed_follows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feed_follows
}

func databasePostToPost (dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		ID: dbPost.ID,
		CreatedAt: dbPost.CreatedAt,
		UpdatedAt: dbPost.PublishedAt,
		Title: dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url: dbPost.Url,
		FeedID: dbPost.FeedID,
	}
}

func databasePostsToPosts (dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, post := range dbPosts {
		posts = append(posts, databasePostToPost(post))
	}
	return posts
}