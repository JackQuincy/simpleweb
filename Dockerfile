FROM scratch

# Go simple web server
ADD main /
 
EXPOSE 80 

RUN chmod +rx /main

CMD ["/main"]
