FROM golang

# Set working dir
WORKDIR /ssr

# Package Manager apt update, upgrade and install Node.js for tailwindcss
RUN apt update -y\
&& apt upgrade -y\
&& curl -fsSL https://deb.nodesource.com/setup_18.x | bash -\
&& apt install nodejs

COPY . .    

# Install Node dependencies and compile tailwind. Install templ. Generate Dummy Data
RUN npm install\
&& npm install -D\
&& go mod download\
&& go install github.com/a-h/templ/cmd/templ@latest\
&& templ generate\
&& go run main.go style

EXPOSE 8080
CMD [ "start" ]
ENTRYPOINT [ "go","run","main.go" ]