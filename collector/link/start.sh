docker build -t cgp-link .
docker stop cgp-link
docker rm cgp-link
docker run -d --name cgp-link cgp-link
