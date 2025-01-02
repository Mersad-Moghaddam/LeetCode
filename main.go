package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<head>
	<title>LeetCode Solutions</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f9;
			color: #333;
			margin: 0;
			padding: 0;
		}
		.container {
			width: 80%%;
			margin: auto;
			overflow: hidden;
		}
		header {
			background: #35424a;
			color: #ffffff;
			padding-top: 30px;
			min-height: 70px;
			border-bottom: #e8491d 3px solid;
		}
		header a {
			color: #ffffff;
			text-decoration: none;
			text-transform: uppercase;
			font-size: 16px;
		}
		header ul {
			padding: 0;
			list-style: none;
		}
		header li {
			float: left;
			display: inline;
			padding: 0 20px 0 20px;
		}
		header #branding {
			float: left;
		}
		header #branding h1 {
			margin: 0;
		}
		header nav {
			float: right;
			margin-top: 10px;
		}
		.showcase {
			min-height: 400px;
			background: url('https://www.example.com/leetcode-banner.jpg') no-repeat 0 -400px;
			text-align: center;
			color: #000000; /* Text color changed to black */
		}
		.showcase h1 {
			margin-top: 100px;
			font-size: 55px;
			margin-bottom: 10px;
		}
		.showcase p {
			font-size: 20px;
		}
		.content {
			padding: 20px;
			background: #ffffff;
		}
		footer {
			background: #35424a;
			color: #ffffff;
			text-align: center;
			padding: 30px 0;
			margin-top: 30px;
		}
	</style>
</head>
<body>
	<header>
		<div class="container">
			<div id="branding">
				<h1>LeetCode Solutions</h1>
			</div>
			<nav>
				<ul>
					<li><a href="/">Home</a></li>
					<li><a href="/about">About</a></li>
					<li><a href="/contact">Contact</a></li>
				</ul>
			</nav>
		</div>
	</header>
	<section class="showcase">
		<div class="container">
			<h1>Welcome to My LeetCode Solutions</h1>
			<p>Tracking my progress and improving my problem-solving skills.</p>
		</div>
	</section>
	<section class="content container">
		<h2>About This Project</h2>
		<p>This repository contains my solutions to various coding challenges from LeetCode. It includes problems from different difficulty levels (Easy, Medium, and Hard), solved using a variety of algorithms and data structures. The goal of this repository is to track my progress as I work through LeetCode problems and improve my problem-solving skills.</p>
		<p>Each solution is implemented in Go, a statically typed, compiled programming language designed at Google. Go is known for its simplicity, efficiency, and strong support for concurrent programming, making it an excellent choice for solving algorithmic problems.</p>
		<p>In this project, you will find solutions to problems involving:</p>
		<ul>
			<li>Arrays and Strings</li>
			<li>Linked Lists</li>
			<li>Stacks and Queues</li>
			<li>Trees and Graphs</li>
			<li>Dynamic Programming</li>
			<li>Sorting and Searching</li>
			<li>Backtracking</li>
			<li>Bit Manipulation</li>
			<li>Math and Geometry</li>
		</ul>
		<p>Each solution is accompanied by a detailed explanation of the approach used, the time and space complexity analysis, and alternative solutions where applicable. The repository is structured in a way that makes it easy to navigate and find solutions to specific problems.</p>
		<p>Whether you are a beginner looking to improve your coding skills or an experienced programmer preparing for technical interviews, this repository offers valuable resources and insights to help you succeed.</p>
		<p>Feel free to explore the solutions, provide feedback, and contribute to the project. Happy coding!</p>
	</section>
	<footer>
		<p>LeetCode Solutions &copy; 2023</p>
	</footer>
</body>
</html>`)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<head>
	<title>About Me - Mersad Moghaddam</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f9;
			color: #333;
			margin: 0;
			padding: 0;
		}
		.container {
			width: 80%%;
			margin: auto;
			overflow: hidden;
		}
		header {
			background: #35424a;
			color: #ffffff;
			padding-top: 30px;
			min-height: 70px;
			border-bottom: #e8491d 3px solid;
		}
		header a {
			color: #ffffff;
			text-decoration: none;
			text-transform: uppercase;
			font-size: 16px;
		}
		header ul {
			padding: 0;
			list-style: none;
		}
		header li {
			float: left;
			display: inline;
			padding: 0 20px 0 20px;
		}
		header #branding {
			float: left;
		}
		header #branding h1 {
			margin: 0;
			color: #ffffff; /* Title color changed to white */
		}
		header nav {
			float: right;
			margin-top: 10px;
		}
		.content {
			padding: 20px;
			background: #ffffff;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			margin-top: 20px;
		}
		footer {
			background: #35424a;
			color: #ffffff;
			text-align: center;
			padding: 30px 0;
			margin-top: 30px;
		}
	</style>
</head>
<body>
	<header>
		<div class="container">
			<div id="branding">
				<h1>About Me</h1>
			</div>
			<nav>
				<ul>
					<li><a href="/">Home</a></li>
					<li><a href="/about">About</a></li>
					<li><a href="/contact">Contact</a></li>
				</ul>
			</nav>
		</div>
	</header>
	<section class="content container">
		<h2>About Me</h2>
		<p>I’m Mersad Moghaddam, a 23-year-old Computer Engineering student from Iran, currently pursuing my degree at Sajad Industrial University. Alongside my studies, I work as a software developer specializing in Flutter, with two years of experience creating cross-platform mobile applications. I am also an active member of my university’s scientific association and a proud participant in the SPACE (ICPC) team.</p>
		<p>My ultimate career goal is to become a Chief Technology Officer (CTO) and lead innovative projects that harness emerging technologies to create impactful solutions. I’m deeply passionate about backend development, scalable software design, and distributed systems. My technical expertise includes Python, C++, Dart, and Go (Golang), with a focus on creating efficient APIs and scalable systems. I also have experience with .NET, relational database design, and implementing robust software architectures.</p>
		<p>I’m particularly fascinated by blockchain technology and its transformative potential across industries like finance and healthcare. Despite financial challenges and economic sanctions, I’ve cultivated resilience and resourcefulness, seeking out affordable or free learning resources to further my knowledge. My GitHub (Mersad-Moghaddam) showcases projects like my Book Library and ongoing e-commerce ventures, including a website for buying and selling saffron built with Go.</p>
		<p>I value clean code, effective software design principles, and user-centric experiences. My work emphasizes SOLID principles, design patterns, and modularity for scalability. Beyond technical skills, I prioritize continuous learning and professional growth. I actively study 20 new English words daily to enhance my language skills and have recently started learning German. My interests also include exploring cybersecurity using tools like Kali Linux and implementing Go-based solutions on this platform.</p>
		<p>As a problem solver, I tackle challenges with patience, creativity, and a results-driven mindset. I balance work, studies, and personal projects with structured time management. My curiosity drives me to stay updated on emerging technologies like serverless computing, AI, and edge computing. I gather insights from platforms like Hacker News to keep my knowledge current and relevant.</p>
		<p>I envision a future where I lead groundbreaking tech initiatives, fostering innovation and collaboration within my teams. For me, success is not just about achievements but also about the impact I create and the legacy I leave behind. My journey is fueled by ambition, learning, and the drive to make a meaningful difference in the world of technology.</p>
	</section>
	<footer>
		<p>LeetCode Solutions &copy; 2023</p>
	</footer>
</body>
</html>`)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<head>
	<title>Contact - Mersad Moghaddam</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f9;
			color: #333;
			margin: 0;
			padding: 0;
		}
		.container {
			width: 80%%;
			margin: auto;
			overflow: hidden;
		}
		header {
			background: #35424a;
			color: #ffffff;
			padding-top: 30px;
			min-height: 70px;
			border-bottom: #e8491d 3px solid;
		}
		header a {
			color: #ffffff;
			text-decoration: none;
			text-transform: uppercase;
			font-size: 16px;
		}
		header ul {
			padding: 0;
			list-style: none;
		}
		header li {
			float: left;
			display: inline;
			padding: 0 20px 0 20px;
		}
		header #branding {
			float: left;
		}
		header #branding h1 {
			margin: 0;
			color: #ffffff; /* Title color changed to white */
		}
		header nav {
			float: right;
			margin-top: 10px;
		}
		.content {
			padding: 20px;
			background: #ffffff;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			margin-top: 20px;
		}
		footer {
			background: #35424a;
			color: #ffffff;
			text-align: center;
			padding: 30px 0;
			margin-top: 30px;
		}
	</style>
</head>
<body>
	<header>
		<div class="container">
			<div id="branding">
				<h1>Contact</h1>
			</div>
			<nav>
				<ul>
					<li><a href="/">Home</a></li>
					<li><a href="/about">About</a></li>
					<li><a href="/contact">Contact</a></li>
				</ul>
			</nav>
		</div>
	</header>
	<section class="content container">
		<h2>Contact Information</h2>
		<p>Mashhad, Iran</p>
		<p>Email: mersad.moghaddam@example.com</p>
		<p>Phone: +98 912 345 6789</p>
	</section>
	<footer>
		<p>LeetCode Solutions &copy; 2023</p>
	</footer>
</body>
</html>`)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
