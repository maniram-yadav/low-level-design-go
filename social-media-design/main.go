package main

import "fmt"

func main() {
	socialMedia := ActivityManagerFacade()

	// Simulate user actions

	user1 := NewUser(1, "John Doe", "john.doe@example.com", "password123")
	user2 := NewUser(2, "Jane Smith", "jane.smith@example.com", "secret123")

	socialMedia.AddUser(user1)
	socialMedia.AddUser(user2)

	//Create posts and add them to the social media platform
	post1 := NewPost("Today, I went to the park and walked 1000 steps!", 1, 1, nil)
	post2 := NewPost("I painted a beautiful sunset on the wall!", 2, 2, nil)
	post3 := NewPost("I've been working on a new video game!", 3, 1, nil)

	err := socialMedia.AddPost(post1)
	if err != nil {
		fmt.Println("\nError adding post:", err)
	}
	err = socialMedia.AddPost(post2)
	if err != nil {
		fmt.Println("Error adding post:", err)
	}
	err = socialMedia.AddPost(post3)
	if err != nil {
		fmt.Println("Error adding post:", err)
	}

	// Get user feeds
	err = getUserFeed(socialMedia, user1.Id)
	if err != nil {
		fmt.Println("\nError getting user feed:", err)
	}
	err = getUserFeed(socialMedia, user2.Id)
	if err != nil {
		fmt.Println("\nError getting user feed:", err)
	}

	// Send friend requests and accept them
	err = socialMedia.SendFriendRequest(1, 2)
	if err != nil {
		fmt.Println("\nError sending friend request:", err)
	}
	err = socialMedia.AcceptFriendRequest(1, 2)
	if err != nil {
		fmt.Println("\nError accepting friend request:", err)
	}

	err = getUserFeed(socialMedia, user1.Id)
	if err != nil {
		fmt.Println("\nError getting user feed:", err)
	}

	// Publish and unpublish posts
	_ = socialMedia.UnPublishPost(1)

	err = getUserFeed(socialMedia, user2.Id)
	if err != nil {
		fmt.Println("Error getting user feed:", err)
	}

	// Comment on a post
	err = socialMedia.CommentPost(1, 2, "I really like this post!")
	if err != nil {
		fmt.Println("Error commenting on post:", err)
	}

	comments := post2.GetComments()
	for _, comment := range comments {
		fmt.Printf("\nComment Detail User %s: %s\n", comment.User.Name, comment.Content)
	}

	// Like a post
	err = socialMedia.LikePost(2, 3)
	if err != nil {
		fmt.Println("Error liking post:", err)
	}

	fmt.Printf("\nPost %d's likes: %d\n", post3.Id, post3.Id)

	err = socialMedia.LikePost(2, 3)
	if err != nil {
		fmt.Println("Error liking post:", err)
	}

	fmt.Printf("Post %d's likes: %d\n", post3.Id, post3.GetLikes())
}

func getUserFeed(socialMedia *ActivityManager, userID int) error {
	feed, err := socialMedia.GetUserFeed(userID)
	if err != nil {
		fmt.Println("Error getting feed posts:", err)
	}
	for _, post := range feed {
		fmt.Printf("\nUser %d Post %d: Post Content : %s\n", post.UserId, post.Id, post.Content)
	}

	return nil
}
