# challenge-intern-bigblue

Hello, get yourself ready for this intern challenge ! :fire:

This README will present you the subject, the code, how to run it and the most important: your missions. 

## The subject

You will work on an e-shop with a list of 6 products (with their name, price, image and "Add to cart" button) fetched from the backend and displayed in the frontend, along with a cart in which you can see items added to cart. 

<img width="1438" alt="image" src="https://user-images.githubusercontent.com/48725727/118649510-de4d5000-b7e3-11eb-8da9-9c298d05402d.png">

## The code

This *intern* folder contains a *backend* folder with incomplete code and a *frontend* folder that contains a src folder with 2 files `index.html` and `index.tsx` and 2 folders *Cart* and *Shop*. 

`index.html` is the HTML page that will be displayed on your browser when you run the code. But there is not much code in it. Indeed the "interesting" code is in the Javascript files whose root is the file `index.jsx` that imports the components `Cart` and `Shop`. 

Running this code displays a page with the e-shop. 

The style is handled through the Bootstrap library (see [Bootstrap docs](https://getbootstrap.com/docs/5.0/getting-started/introduction/)). So the style is determined by all the `className=...`. You don't have to take care of that except if you want to reuse some style for what you add in the code. 


## How to run it

### Backend 

⚠️ In order to run the backend you first need to complete the backend code. For that search for the `TODO` lines in the `products.go` file and write the missing lines. Don't worry you don't need to add many code lines. 

Once completed you can run the backend: to start the server, open a new terminal, go to `intern/backend` and run the following command: `go run products.go`. 

### Frontend

In order to run the frontend part of this exercise you need to get Node.js and npm. You can get both by downloading Node.js [here](https://nodejs.org/en/download/) and following the installation instructions. 
You can check that you have correctly installed Node.js and npm by running in your command terminal: 
```
node -v
npm -v
```

If npm was not installed, try installing it with: 
```
npm install -g npm
```

Then download this repository (by downloading the .zip and extracting it for example), go to the `intern/frontend/` and run 
```
npm install
```
to install all the necessary packages. Finally you can run the code with the command
```
npm start
```
Frontend is now running and you can start coding!


## Your missions

Backend missions: 
- Complete the backend code where there are TODOs (if not already done):
  - Add code to get the list of products in the Handler. 
  - Add code to start a local web server at port 8080 that handles all requests with Handler in the main function. 

Frontend missions: 
- Compute the total price of the products in the cart and display it (fill the code line 25 in `Cart.jsx`)
- Add a remove button next to each product in the cart.
- Show remaining stock for each product in the shop.
- Prevent user from adding a product to cart if there is no stock available anymore. 
- Don't display an item line from the cart if the quantity is zero. 

You are free to do whatever you want concerning the design of what you add but try to make it look good. You can reuse the style of already present components. 

You must submit the project as a git repository (github.com, bitbucket.com, gitlab.com). Repository must avoid containing the words bigblue and challenge. 

Have fun :rocket: :wink:
