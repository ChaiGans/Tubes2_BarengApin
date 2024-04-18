"use client";
import React, { useState, useEffect } from "react";

interface Suggestion {
	title: string;
	link: string;
}

const SearchInput: React.FC = () => {
	const [searchTerm, setSearchTerm] = useState<string>("");
	const [suggestions, setSuggestions] = useState<Suggestion[]>([]);
	const [isFetching, setIsFetching] = useState<boolean>(false);
	const [apiResponse, setApiResponse] = useState<any>(null); // State to store the raw API response

	useEffect(() => {
		const timerId = setTimeout(() => {
			if (searchTerm) {
				setIsFetching(true);
				fetchSuggestions(searchTerm).then((suggestions) => {
					setSuggestions(suggestions);
					setIsFetching(false);
				});
			} else {
				setSuggestions([]);
				setApiResponse(null); // Clear the API response when there is no search term
			}
		}, 300); // Debounce time is 300 ms

		return () => clearTimeout(timerId);
	}, [searchTerm]);

	const fetchSuggestions = async (
		searchTerm: string
	): Promise<Suggestion[]> => {
		const url = new URL("https://id.wikipedia.org/w/api.php");
		url.searchParams.append("action", "opensearch");
		url.searchParams.append("search", searchTerm);
		url.searchParams.append("format", "json");
		url.searchParams.append("limit", "5"); // Specify the maximum number of search results
		url.searchParams.append("origin", "*"); // To handle CORS

		// console.log(url.href);

		try {
			const response = await fetch(url.href);
			const data = await response.json();
			setApiResponse(data); // Store the raw API response in state

			// The response format for "opensearch" is an array with four elements:
			const titles = data[1];
			const links = data[3];

			return titles.map((title: string, index: number) => ({
				title: title,
				link: links[index],
			}));
		} catch (error) {
			console.error("Failed to fetch data:", error);
			return []; // Return an empty array on error
		}
	};

	return (
		<div>
			<input
				type="text"
				value={searchTerm}
				onChange={(e) => setSearchTerm(e.target.value)}
				placeholder="Search Wikipedia"
			/>
			{isFetching ? (
				<div>Loading...</div>
			) : (
				<ul>
					{suggestions.map((s, index) => (
						<li key={index}>
							<a href={s.link} target="_blank" rel="noopener noreferrer">
								{s.title}
							</a>
						</li>
					))}
				</ul>
			)}
			{/* Optionally display the raw JSON response for debugging */}
			<div>
				<h3>API Response:</h3>
				<pre>{JSON.stringify(apiResponse, null, 2)}</pre>
			</div>
		</div>
	);
};

export default SearchInput;
