# This script copies all relevant files from this repository to your server.
# While you could use 'git pull' directly on the server, that approach isn't
# ideal if you want to exclude certain files from version control. For example,
# large folders containing images might not be suitable for GitHub due to
# storage limits. To avoid this, we simply copy the entire project to the
# server.

# TODO: make a scripterino that copies stuff from your computer to your
# production server
# ex: rsync ./ root@yourserver://home/user/myproject/

# I use rsync, cause it's ezpz. You could use scp or whatever you like. The
# important part is to make it easy so you can just run `./build.sh` when you
# want to upload to production.
