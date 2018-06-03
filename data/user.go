package data

import (
	"time"
)

type User struct{
	Id int
	Uuid string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

type Session struct{
	Id int
	Uuid string
	Email string
	UserId int
	CreatedAt time.Time
}

func (user *User) Create() error{
	stmt, err := Db.Prepare("insert into users (uuid, name, email, password, created_at) values($1, $2, $3, $4, $5)" +
		"returning id, uuid, created_at")
	defer stmt.Close()
	if err != nil{
		return err
	}
	uuid, err := createUUID()
	if err != nil{
		return err
	}
	return stmt.QueryRow(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(
		&user.Id, &user.Uuid, &user.CreatedAt)
}

func (user *User) CreateSession() (Session, error){
	stmt, err := Db.Prepare("insert into sessions (uuid, email, user_id, created_at) values($1, $2, $3, $4)" +
		"returning id,uuid,created_at")
	defer stmt.Close()
	sess := Session{
		Email:user.Email,
		UserId:user.Id,
	}
	if err != nil{
		return sess, err
	}
	uuid, err := createUUID()
	if err != nil{
		return sess, err
	}
	err = stmt.QueryRow(uuid, user.Email, user.Id, time.Now()).Scan(&sess.Id, &sess.Uuid, &sess.CreatedAt)
	return sess, err
}

func (user *User) GetSession() (Session, error){
	sess := Session{}
	err := Db.QueryRow("select id,uuid,email,user_id,created_at from sessions where user_id=$1", user.Id).
		Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId, &sess.CreatedAt)
	return sess, err
}

func (user *User) CreateThread(topic string) (Thread, error){
	stmt, err := Db.Prepare("insert into threads (uuid, topic, user_id, created_at) values($1, $2, $3, $4)" +
		"returning id, uuid, topic, user_id, created_at")
	defer stmt.Close()
	t := Thread{}
	if err != nil{
		return t, err
	}
	uuid, err := createUUID()
	if err != nil{
		return t, err
	}
	err = stmt.QueryRow(uuid, topic, user.Id, time.Now()).Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreatedAt)
	return t, err
}

func (user *User) CreatePost(t Thread, body string) (Post, error){
	stmt, err := Db.Prepare("insert into posts (uuid, body, user_id, thread_id, created_at) values($1, $2, $3, $4, $5)" +
		"returning id, uuid, body, user_id, thread_id, created_at")
	defer stmt.Close()
	post := Post{}
	if err != nil{
		return post, err
	}

	uuid, err := createUUID()
	if err != nil{
		return post, err
	}
	err = stmt.QueryRow(uuid, body, user.Id, t.Id, time.Now()).
		Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return post, err
}

func GetUserByEmail(email string) (User, error){
	user := User{}
	err := Db.QueryRow("select id, uuid, name, email, password, created_at from users where email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

func (sess * Session) Check() (bool, error){
	err := Db.QueryRow("select id, uuid, email, user_id, created_at from sessions where uuid=$1", sess.Uuid).
		Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId, &sess.CreatedAt)
	if err != nil{
		return false, err
	}
	if sess.Id!=0{
		return true, nil
	}
	return false, nil
}

func (sess * Session) DeleteByUUID() error{
	stmt, err := Db.Prepare("delete from sessions where uuid=$1")
	defer stmt.Close()
	if err != nil{
		return err
	}
	_, err = stmt.Exec(sess.Uuid)
	return err
}

func (sess * Session) GetUser() (User, error){
	user := User{}
	err := Db.QueryRow("select id, uuid, name, email, created_at from users where id=$1", sess.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return user, err
}