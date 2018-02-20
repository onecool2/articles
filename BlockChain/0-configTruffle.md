#1. Install truffle

    npm install -g truffle 3.2.1
    
#2. Install testRPC 

    npm install -g ethereumjs-testrpc
    
#3. Truffle init

    truffle init webpack
    
#4. Truffle compile

    truffle compile
    
#5. Truffle migrate

    npm run dev
    
#6. then you can access it from localhost:8080

# Webpack Truffle Box

This box it our most bare official implementation with Webpack. Includes contracts, migrations, tests, user interface and webpack build pipeline.

## Installation

1. Install Truffle globally.
    ```javascript
    npm install -g truffle
    ```

2. Download the box. This also takes care of installing the necessary dependencies.
    ```javascript
    truffle unbox webpack
    ```

3. Run the development console.
    ```javascript
    truffle develop
    ```

4. Compile and migrate the smart contracts. Note inside the development console we don't preface commands with `truffle`.
    ```javascript
    compile
    migrate
    ```

5. Run the webpack server for front-end hot reloading (outside the development console). Smart contract changes must be manually recompiled and migrated.
    ```javascript
    // Serves the front-end on http://localhost:8080
    npm run dev
    ```

6. Truffle can run tests written in Solidity or JavaScript against your smart contracts. Note the command varies slightly if you're in or outside of the development console.
  ```javascript
  // If inside the development console.
  test

  // If outside the development console..
  truffle test
  ```

