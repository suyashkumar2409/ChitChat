package data

import "time"

type Thread struct{
	Id int
	Uuid string
	Topic string
	UserId int
	CreatedAt time.Time
}

type Post struct{
	Id int
	Uuid string
	Body string
	UserId int
	ThreadId int
	CreatedAt time.Time
}

func GetThreads() (threads []Thread, err error){
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	defer rows.Close()
	if err != nil{
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil{
			return
		}
		threads = append(threads, th)
	}
	return
}

func GetThreadByUUID(uuid string) (Thread,error){
	t := Thread{}
	err := Db.QueryRow("select id, uuid, topic, user_id, created_at from threads where uuid=$1", uuid).
		Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreatedAt)
	return t, err
}

func (t *Thread) NumReplies() (count int){
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", t.Id)
	defer rows.Close()
	if err != nil{
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil{
			return
		}
	}
	return
}

func (t* Thread) CreatedAtDate() string{
	return formatTimeToString(t.CreatedAt)
}

func (t * Thread) GetPosts() ([]Post, error){
	rows, err := Db.Query("select id, uuid, body, user_id, thread_id, created_at from posts where thread_id=$1", t.Id)
	defer rows.Close()
	if err != nil{
		return nil, err
	}
	posts := make([]Post,0)
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
		if err != nil{
			return posts, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (t * Thread) GetUser() (User){
	user := User{}
	Db.QueryRow("select id, uuid, name, email, created_at from users where id=$1", t.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return user
}

func (p * Post) CreatedAtDate() (string){
	return formatTimeToString(p.CreatedAt)
}

func (p * Post) GetUser() (User){
	user := User{}
	Db.QueryRow("select id, uuid, name, email, created_at from users where id=$1", p.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return user
}