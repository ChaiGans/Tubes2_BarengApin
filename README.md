<h1 align="center">Tugas Besar 2 IF2211 Strategy Algorithm</h1>
<h1 align="center">Group BarengApin</h3>
<h3 align="center">Utilization of IDS Algorithm and BFS in the WikiRace Game</p>

## Table of Contents

- [Overview](#overview)
- [Abstraction](#abstraction)
- [Built With](#built-with)
- [Installation With Docker](#installation-with-docker)
- [Installation Without Docker](#installation-without-docker)
- [Links](#links)

## Overview

![messageImage_1714108706482](https://github.com/ChaiGans/Tubes2_BarengApin/assets/113753352/453096df-c257-49e0-8908-b5c0fb78b91f)
![messageImage_1714107805613](https://github.com/ChaiGans/Tubes2_BarengApin/assets/113753352/9535b50a-8b2f-4b03-83a8-ccb9209e47a2)
![messageImage_1714107894930](https://github.com/ChaiGans/Tubes2_BarengApin/assets/113753352/b3b19bc4-93ab-4db1-aa80-29bca324417a)

Our Team members :

- 13522019 - Wilson Yusda
- 13522045 - Elbert Chailes
- 13522081 - Albert

<p>Our Lecturer : Dr. Ir. Rinaldi Munir, M.T.</p>

Here is the purpose of making this project :

- To fulfill the requirements of the second project assignment for the IF2211 Strategy Algorithm course.
- To implement IDS and BFS search algorithms to find the shortest path from one URL to another.
- To provide developers with the opportunity to learn web scraping using Go-colly to extract links from Wikipedia URLs.
- To serve as a platform for developers to explore website development, including the creation of algorithms and logic using the Go programming language.
- To contribute to technological advancement, particularly in the field of optimizing pathfinding.

## Abstraction

In this project, the developer was inspired by a game called Wikirace, which has gained popularity across the internet. The developer aimed to learn BFS and IDS algorithms for pathfinding and web scraping, with the goal of finding the shortest path from one URL to another. Overall, the website features four main functionalities for finding the shortest path between two nodes. This allows the developer to analyze and compare which algorithm is better suited for different cases—determining when BFS or IDS performs more effectively. Additionally, the developer included a feature that enables users to search for multiple paths, while still prioritizing the shortest path, or paths with the same depth as the initially found shortest path. Essentially, the project revolves around finding the shortest path, with the starting link referred to as the "starting node," and the goal link as the "goal node." To achieve this, the developer utilized web scraping to extract URLs using Go-colly, a library in Golang. Each extracted link is treated as a node, resulting in a graph organization that aligns with the characteristics of BFS and IDS algorithms.

## Problem Solving Steps with BFS Algorithm

1. Receive the initial URL from the frontend, treated as the starting node for the search.
2. Scrape the starting URL to get related child URLs, considered as nodes to explore further, and place them in a queue for search.
3. Evaluate each scraped URL. If it doesn't match the target URL, enqueue it again for further exploration at a deeper depth.
4. The search process continues, checking each scraped URL for a match with the target URL. If found, the algorithm stops and returns the URL as the search result.
5. If the target URL is not found in the first iteration, the queue generated from the previous step is iterated for further scraping and checking.
6. Scraping and checking steps are repeated continuously until the target is found or all possibilities have been explored.

## Problem Solving Steps with IDS Algorithm

1. Receive the initial and target URLs for the search, treated as the starting and goal nodes, respectively.
2. Scrape the starting node and generate child nodes, consisting of links found on the parent node. Each generation adds one depth level.
3. Check if the generated child node contains the target URL. There are two possible outcomes:
4. If the child node is the goal node, the algorithm stops and returns the result.
5. If the child node is not the goal node, the search continues to explore sibling nodes at the same depth level.
6. Child generation (steps 2 and 3) continues repeatedly with a limited depth, such as depth n (DLS with depth n).
7. IDS iterates child generation, progressively increasing depth, until a solution is found by repeatedly performing DLS with incremental depth levels.

## Key Differences between BFS and IDS Algorithms

1. BFS focuses on finding the first valid path from the starting node to the target. It stops the search immediately upon finding the first match, making it efficient when only one solution is needed.
2. IDS aims to identify all possible paths from the starting node to the target at the same depth level. It continues the search to gather all potential paths at the same depth, providing a comprehensive view but requiring more resources and time to execute.

## Built With

- [NextJS](https://nextjs.org/docs)
- [GO Lang](https://go.dev/)
- [Tailwind](https://tailwindcss.com/)
- [Gin](https://gin-gonic.com/docs/)

## Prerequisites

To run this project, you will need to perform several installations, including:

- `HTML5` : This is the markup language used for structuring the content of web pages. It's a fundamental part of web development and is typically assumed to be available in web development environments.
- `Node.js` : Node.js is essential for running JavaScript on the server-side and for managing JavaScript-based build processes, including those used in React applications.
- `npm` (Node package manager) : npm is indeed the package manager for JavaScript and is used to install and manage JavaScript packages and libraries, including those required for React development.
- `Go language version 1.18 or higher` : This is necessary if your project involves server-side development using the Go programming language. Go is used for building the backend of web applications.
- `Docker`: This is platform that allows developers to package and distribute applications and their dependencies in isolated containers.

**NOTE : It is a must to have docker to run this project. So, download [Docker](https://www.docker.com/products/docker-desktop/) in the internet first, before continuing to the program**

## Installation With Docker

Make sure Docker is installed on your system. If not, install [Docker](https://www.docker.com/products/docker-desktop/) according to the official instructions from the Docker website. If Docker is already installed, follow these steps:

1. Clone this repository :

```shell
git clone https://github.com/ChaiGans/Tubes2_BarengApin.git
```

2. Navigate to the root directory of the program by running the following command in the terminal:

```shell
cd ./Tubes2_BarengApin
```

3. Ensure Docker Desktop is running. Once the user is in the root directory, run the following command in the terminal:

```shell
docker compose build
```

4. Once the docker compose build command has finished, run the following command:

```shell
docker compose up
```

5. To access the website, go to the following link in your web browser: [http://localhost:3000](http://localhost:3000)
6. After successfully launching the website, users can choose the search algorithm, either using BFS or IDS. Once the user selects the search algorithm, they need to enter the title of the source and target Wikipedia articles. The program will also provide Wikipedia article recommendations based on the titles entered by the user. The program will display the shortest route between the two articles in the form of a graph visualization, along with execution time, the number of articles traversed, and the search depth.

## Installation Without Docker

Follow these steps:

1. Clone this repository :

```shell
git clone https://github.com/ChaiGans/Tubes2_BarengApin.git
```

2. Navigate to the src directory of the program by running the following command in the terminal:

```shell
cd ./Tubes2_BarengApin/src
```

3. Open another terminal in the same folder directory. Now, you should have two terminals open, both in the Tubes2_BarengApin/src directory.

First Terminal (To run the frontend server)

1. Change directory to the frontend folder and execute the following commands

```shell
cd frontend
npm install
npm run start
```

2. Open a web browser and navigate to http://127.0.0.1:3000. Note: The port number may vary if the default port is in use. Check the terminal to determine which port the server is running on.

Second Terminal (To run the backend server)

1. Change directory to the backend folder and execute the following commands

```shell
cd backend
```

2. Install the required packages:

```shell
go get .
```

3. Start the Backend server:

```shell
go run .
```

The backend server should now be running at http://127.0.0.1:8080.

4. After successfully launching the website, users can choose the search algorithm, either using BFS or IDS. Once the user selects the search algorithm, they need to enter the title of the source and target Wikipedia articles. The program will also provide Wikipedia article recommendations based on the titles entered by the user. The program will display the shortest route between the two articles in the form of a graph visualization, along with execution time, the number of articles traversed, and the search depth.

## Links

- Repository : [https://github.com/ChaiGans/Tubes2_BarengApin/](https://github.com/ChaiGans/Tubes2_BarengApin/)
