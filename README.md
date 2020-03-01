# Publish it everywhere Mattermost Bot

Publish it everywhere is a mattermost bot to publish your content across multiple social media platform with a single command without leaving mattermost.

Typically, in a company, when a new feature is added to their product, they post it across multiple social media platform(Twitter, LinkedIn, Facebook etc.). Mostly the content they post are same, and for that they have to login to every social media, and post it everywhere. This is a tedious task. `Publish-it-everywhere` saves your precious time by posting your content to multiple social media platform with a single command in mattermost.

One more use case would be, If in your company, you want  someone to post content on behalf of your company, but you don't want to share the credentials with that person. In that case, you can just create a private channel in mattermost, add the relavant people and PIE bot there. All the people in that channel will be able to post the content on behalf of your company using PIE bot.

## Installation

1. Follow this page(https://developers.mattermost.com/contribute/getting-started/) to setup `mattermost-server` and `mattermost-webapp`.

2. Create your account and a team from the UI.

3. Create a Bot Account:

   1. Get into the mattermost-server directory:

      ```bash
      cd mattermost-server
      ```

   2. Login to your account from mmctl into the teamname you created earlier.(You can list all the available teams by using `./bin/mmctl team list`)

      ```bash
      ./bin/mmctl auth login http://localhost:8065 --name <teamname> --username <your_username> --password <your_password>
      ```

   3. Create an account for bot.

      ```bash
      ./bin/mmctl user create --email="<bot@email.address>" --password="<bot_password>" --username="<bot_username>"
      ```

   4. Add the bot to the team:

      ```bash
      ./bin/mmctl team users add <teamname> <bot_username>
      ```

4. Run the Server

   1. Clone this repository: 

      ```bash
      git clone git@github.com:shibasisp/publish_it_everywhere.git
      ```

   2. Create a file called `.env`. Copy the content from `.env.template` to `.env`

   3. Fill the actual credentials in the quotes(`DBNAME`, `DBURL`, `TWITTER_CONSUMER_KEY`, `TWITTER_CONSUMER_SECRET`,`LINKEDIN_CLIENT_ID`,`LINKEDIN_CLIENT_SECRET` and `SELF_URL`) to the `.env` file
      
      - `SELF_URL` is the domain name of the server.
      - `DBNAME`, `DBURL` and `SELF_URL` are optional. Default values will be used if no value is specified.
      - `TWITTER_CONSUMER_KEY`, `TWITTER_CONSUMER_SECRET`,`LINKEDIN_CLIENT_ID`,`LINKEDIN_CLIENT_SECRET` needs to be get from the Twitter app and LinkedIn app respectively. More details on how to create Twitter app and LinkedIn app can be found [here](https://developer.twitter.com/apps) and [here](https://docs.microsoft.com/en-us/linkedin/shared/authentication/authorization-code-flow) respectively.
      

   4. Export all the environment variables defined to the shell.

      ```
      export .env
      ```

   5. To use go modules, you need to export  `GO111MODULE=on`

      ```
      export GO111MODULE=on 
      ```

   6. Get all the dependencies

      ```
      go get
      ```

   7. Run the server

      ```bash
      go run main.go
      ```

5. Run the bot

   1. Clone this publish-it-everywhere-bot repository: 

      ```bash
      git clone git@github.com:shibasisp/publish-it-everywhere-bot.git
      ```

   2. Start the bot

      ```bash
      make run
      ```

## Demo

Video Link: https://vimeo.com/394777943

<p align="center">
  <img src="https://github.com/shibasisp/publish-it-everywhere-bot/raw/master/demo/demo.gif" />
</p>

## Usage

First of all you need to authorize Publish_it_everywhere bot to publish to twitter and LinkedIn. For that authenticate will give you a link. You need to click on the link and authorize the app.

1. Authenticate with LinkedIn:

   ```
   :authenticate_linkedin
   ```

2. Authenticate with Twitter:

   ```
   :authenticate_twitter
   ```

3. Publish to Twitter:

   ```
   :publish_twitter <YOUR POST>
   ```

4. Publish to LinkedIn:

   ```
   :publish_linkedin <YOUR POST>
   ```

5. Publish it everywhere:

   ```
   :publish_it_everywhere <YOUR POST>
   ```
   
## Team Members
- Aahel Guha(https://github.com/aahel)
- Asutosh Sahoo(https://github.com/asutosh97)
- Prateek Patnaik(https://github.com/Northen-Light)
- Shibasis Patel(https://github.com/shibasisp)
