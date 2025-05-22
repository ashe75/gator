#GATOR

##Description
Gator allows to follow and keep track on RSS feeds.

##Installation
To use gator you will need Go runtim and Postgres installed

For installation you can use
`go install github.com/ashe75/gator`

You need to set up config file in your home repository wit the name ".gatorconfig.json". Example:
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable","current_user_name":"bob"}

Then you can run by `./gator command args`

##Commands
login $username - login user
register $username - add new user
reset - reset all data
users - list all users
agg $interval - add posts to database with this interval
addfeed $name $url - add new feed
feeds - list all feeds
follow $url - follow feed
unfollow $url - unfollow feed
following - list all followed feeds
browse $limit - browse $limit amount of posts
