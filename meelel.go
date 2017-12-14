/*
   Copyright 2017 OÜ Nevermore

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package meelel

import (
	"context"
	"errors"
	"time"

	"google.golang.org/appengine/datastore"

	"github.com/mjibson/goon"
)

type Meelel struct {
	g *goon.Goon
}

type Post struct {
	_kind           string    `goon:"kind,MeelelPost"`
	Id              string    `datastore:"-" goon:"id"`
	Active          bool      `datastore:"a"`
	Created         time.Time `datastore:"c"`
	Modified        time.Time `datastore:"m"`
	Title           string    `datastore:"t,noindex"`
	PageTitle       string    `datastore:"mt,noindex"`
	MetaDescription string    `datastore:"md,noindex"`
	Content         string    `datastore:"co,noindex`
}

func New(c context.Context) *Meelel {
	return NewWithGoon(goon.FromContext(c))
}

func NewWithGoon(g *goon.Goon) *Meelel {
	return &Meelel{g: g}
}

func (m Meelel) SavePost(postData *Post) (*Post, error) {
	if postData.Id == "" {
		return nil, errors.New("No post id provided!")
	}
	post := &Post{Id: postData.Id}

	err := m.g.RunInTransaction(func(tg *goon.Goon) error {
		now := time.Now()

		// Attempt to get the existing data for this post
		if err := tg.Get(post); err == datastore.ErrNoSuchEntity {
			// Post doesn't exist yet, so set the created time
			post.Created = now
		} else if err != nil {
			return err
		}

		// Update the post data
		post.Active = postData.Active
		post.Title = postData.Title
		post.PageTitle = postData.PageTitle
		post.MetaDescription = postData.MetaDescription
		post.Content = postData.Content

		// Set the modified time
		post.Modified = now

		// Save the updated data
		if _, err := tg.Put(post); err != nil {
			return err
		}

		return nil
	}, &datastore.TransactionOptions{XG: false})

	if err != nil {
		return nil, err
	}
	return post, nil
}

func (m Meelel) GetPost(id string) (*Post, error) {
	post := &Post{Id: id}
	if err := m.g.Get(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (m Meelel) DeletePost(id string) error {
	post := &Post{Id: id}
	if err := m.g.Delete(m.g.Key(post)); err != nil {
		return err
	}
	return nil
}
