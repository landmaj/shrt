shrt
----
Simple link shortener. This is my very first project in Go and so you should not expect too much from it.

### Installation
For an easy deployment, use Dokku. [Let's Encrypt](https://github.com/dokku/dokku-letsencrypt)
and [Postgres](https://github.com/dokku/dokku-postgres) plugins are required.

```
dokku apps:create shrt
dokku domains:add shrt shrt.w-ski.dev
dokku letsencrypt shrt
```

Configuration is stored in environmental variables. They can be tricky in Dokku, so check
[the documentation](http://dokku.viewdocs.io/dokku/configuration/environment-variables/).
All the required variables are listed in [`.env`](.env) file.

```
dokku config:set shrt ROOT_DOMAIN=shrt.w-ski.dev
```

Create and link Postgres database. The second step will create `DATABASE_URL` environmental
variable available to the app.

```
export POSTGRES_IMAGE_VERSION="11.4"
dokku postgres:create shrt_db
dokku postgres:link shrt_db shrt
```

Add your public SSH key to authorized keys on dokku account, add git remote
and push your code.

```
git remote add dokku dokku@dokku_server:shrt
git push dokku master
```

## License

This project is released under the [MIT License](LICENSE).
