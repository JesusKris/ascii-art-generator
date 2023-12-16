
#Builds an image from Dockerfile called auditimage.
docker image build -f Dockerfile -t auditimage .

#Displays all the images on your system.
docker images

#Runs an image called auditimage in a container called auditcontainer in 8080port.
docker container run -p 8080:8080 --detach --name auditcontainer auditimage

#Goes into auditcontainer container root.
docker exec -it auditcontainer /bin/bash


