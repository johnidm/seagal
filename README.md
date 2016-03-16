# seagal

![](https://raw.githubusercontent.com/johnidm/seagal/master/seagal-logo.png)

This is an elegant and simple way to share blog posts in Slack channel with social share buttons.
 
#### How to use
 
The endpoint https://nameless-castle-24251.herokuapp.com/ can be used to test.
 
First of the all you need to create a **Slash Command**, see image:

<center>![](https://raw.githubusercontent.com/johnidm/seagal/master/bot-slack-share-social-post.png)</center>

Now you can share your blog post using the command `/share <URL of blog post>`, example `/share http://www.johnidouglas.com.br/django-migrations-reversible-migrations/ `
 
> Note: Your page should have the graph objects og:title e og:url, for more details access  http://ogp.me/
 
#### Deploy on Heroku 

If you want to install and create your own endpoint, follow this steps.
 
Get the project

```
cd $GOPATH/src/
go get github.com/johnidm/seagal
```
 
Create an app on Heroku
```
heroku login
heroku create
```

Deploy your application
```
git push heroku master
```
 
Ready, you can use a new endpoint.

#### See result

![](https://raw.githubusercontent.com/johnidm/seagal/master/seagal-example.png)
