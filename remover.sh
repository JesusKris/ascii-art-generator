#Stops the container called auditcontainer.
docker stop auditcontainer 

#Deletes the container called auditcontainer.
docker rm auditcontainer 

#Deletes the image called auditimage.
docker image rm auditimage

#Deletes the BaseImage.
docker image rm golang:1.16.5-buster 