/* eslint-disable react/display-name */
"use client";
import React, { useState, useEffect } from "react";
import Image from "next/image";
import reverse from "../../public/swap.png";
import search from "../../public/wiki.png";
import Loading from "./Loading";
import { color } from "d3";
interface Suggestion {
	title: string;
	link: string;
}
interface SearchInputProps {
	onSearchResult: (result: any) => void;
	id: string;
}
const SearchInput: React.FC<SearchInputProps> = ({ onSearchResult, id }) => {
	const [searchTerm1, setSearchTerm1] = useState<string>("");
	const [suggestions1, setSuggestions1] = useState<Suggestion[]>([]);
	const [isFetching1, setIsFetching1] = useState<boolean>(false);
	const [apiResponse, setApiResponse] = useState<any>(null);
	const [showSuggestions1, setShowSuggestions1] = useState<boolean>(true);
	const [showLink1, setLink1] = useState<string>("");
	const [isLoading, setIsLoading] = useState(false);
	// States for the second input
	const [searchTerm2, setSearchTerm2] = useState<string>("");
	const [suggestions2, setSuggestions2] = useState<Suggestion[]>([]);
	const [isFetching2, setIsFetching2] = useState<boolean>(false);
	const [showSuggestions2, setShowSuggestions2] = useState<boolean>(true);
	const [showLink2, setLink2] = useState<string>("");
	const [searchMethod, setSearchMethod] = useState<"BFS" | "IDS">("BFS");
	const [pathing, setPathing] = useState<"SinglePath" | "MultiplePath">(
		"SinglePath"
	);
	const toggleSearchMethod = () => {
		setSearchMethod(searchMethod === "BFS" ? "IDS" : "BFS");
	};
	const togglePathing = () => {
		setPathing(pathing === "SinglePath" ? "MultiplePath" : "SinglePath");
	};
	const handleSuggestionClick = (
		setSearchTerm: React.Dispatch<React.SetStateAction<string>>,
		setLink: React.Dispatch<React.SetStateAction<string>>,
		setShowSuggestions: React.Dispatch<React.SetStateAction<boolean>>,
		suggestion: Suggestion
	) => {
		setSearchTerm(suggestion.title);
		setLink(suggestion.link);
		setShowSuggestions(false);
	};
	const swapSearchTerms = () => {
		setSearchTerm1(searchTerm2);
		setSearchTerm2(searchTerm1);
		setLink1(showLink2);
		setLink2(showLink1);
	};
	useEffect(() => {
		const timerId = setTimeout(() => {
			if (searchTerm1) {
				setIsFetching1(true);
				fetchSuggestions(searchTerm1).then((suggestions) => {
					setSuggestions1(suggestions);
					setIsFetching1(false);
				});
			} else {
				setSuggestions1([]);
			}
		}, 300);

		return () => clearTimeout(timerId);
	}, [searchTerm1]);
	const handleSearchChange1 = (e: React.ChangeEvent<HTMLInputElement>) => {
		const formattedInput = e.target.value.trim().replace(/ /g, '_');
		const words = formattedInput.split('_');
		const capitalizedWords = words.map(word => 
			word.charAt(0).toUpperCase() + word.slice(1)
		);
		const capitalizedInput = capitalizedWords.join('_');
		const url = `https://en.wikipedia.org/wiki/${capitalizedInput}`;
        setSearchTerm1(e.target.value);
        setLink1(url);  
    };
	const handleSearchChange2 = (e: React.ChangeEvent<HTMLInputElement>) => {
		const formattedInput = e.target.value.trim().replace(/ /g, '_');
		const words = formattedInput.split('_');
		const capitalizedWords = words.map(word => 
			word.charAt(0).toUpperCase() + word.slice(1)
		);
		const capitalizedInput = capitalizedWords.join('_');
		const url = `https://en.wikipedia.org/wiki/${capitalizedInput}`;
        setSearchTerm2(e.target.value);
        setLink2(url);  
    };
	useEffect(() => {
		const timerId = setTimeout(() => {
			if (searchTerm2) {
				setIsFetching2(true);
				fetchSuggestions(searchTerm2).then((suggestions) => {
					setSuggestions2(suggestions);
					setIsFetching2(false);
				});
			} else {
				setSuggestions2([]);
			}
		}, 300);

		return () => clearTimeout(timerId);
	}, [searchTerm2]);

	const fetchSuggestions = async (
		searchTerm: string
	): Promise<Suggestion[]> => {
		const url = new URL("https://en.wikipedia.org/w/api.php");
		url.searchParams.append("action", "opensearch");
		url.searchParams.append("search", searchTerm);
		url.searchParams.append("format", "json");
		url.searchParams.append("limit", "5");
		url.searchParams.append("origin", "*");
		try {
			const response = await fetch(url.href);
			const data = await response.json();
			setApiResponse(data);
			const titles = data[1];
			const links = data[3];
			return titles.map((title: string, index: number) => ({
				title: title,
				link: links[index],
			}));
		} catch (error) {
			console.error("Failed to fetch data:", error);
			return [];
		}
	};
const handleSubmit = async () => {
    let algoChoice;
    if (searchMethod === "BFS" && pathing === "SinglePath") {
        algoChoice = 2;
    } else if (searchMethod === "BFS" && pathing === "MultiplePath") {
        algoChoice = 4;
    } else if (searchMethod === "IDS" && pathing === "SinglePath") {
        algoChoice = 1;
    } else if (searchMethod === "IDS" && pathing === "MultiplePath") {
        algoChoice = 3;
    }

    // Function to check if a Wikipedia link is valid
    const validateLink = async (link:string) => {
        const pageTitle = link.split('/wiki/')[1];
        const apiUrl = `https://en.wikipedia.org/api/rest_v1/page/summary/${pageTitle}`;
        try {
            const response = await fetch(apiUrl);
            return response.ok; // Returns true if page exists, false otherwise
        } catch (error) {
            console.error('Error checking link:', link, error);
            return false; // Assume link is invalid if there's an error
        }
    };

    setIsLoading(true);
    try {
        // Validate links
        const isValidLink1 = await validateLink(showLink1);
        const isValidLink2 = await validateLink(showLink2);

        if (!isValidLink1 || !isValidLink2) {
            alert('One or more provided Wikipedia links are invalid.');
            return; // Exit the function if any link is invalid
        }
		if (searchTerm1 == searchTerm2){
			alert('Same Hyperlink detected, Search cancelled');
            return; // Exit the function if any link is invalid
		}
		if (showLink1 == showLink2){
			alert('Same Hyperlink detected, Search cancelled');
            return; // Exit the function if any link is invalid
		}
        const data = {
            StartTitleLink: showLink1,
            GoalTitleLink: showLink2,
            AlgoChoice: algoChoice,
        };

        console.log(data);
        console.log("button clicked");

        const response = await fetch("http://localhost:8080/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });

        console.log("waiting bos");
        if (!response.ok) {
            throw new Error("Network response was not ok.");
        }
        const result = await response.json();
        console.log("uda dapat");
        onSearchResult(result);
    } catch (error) {
        console.error("Failed to fetch data:", error);
    } finally {
        setIsLoading(false); 
    }
};
	return (
		<div className="p-4 ml-20 mr-20" id={id}>
			{isLoading && (
				<div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-50">
					<Loading></Loading>
				</div>
			)}
			<h1 className="text-[#7E5FFF] text-5xl text-center font-black mt-[50px] mb-10">
				Hyperlink Search
			</h1>
			<div className="input1">
				<p className="text-white text-xl font-extrabold mb-2">From</p>
				<input
					type="text"
					value={searchTerm1}
					onChange={handleSearchChange1}
					onFocus={() => setShowSuggestions1(true)}
					className="w-full p-2 mb-4 text-white font-bold bg-[#212122] rounded-md focus:border-[#7E5FFF] focus:border-2 placeholder-[#7E5FFF] focus:outline-none"
					placeholder="Search Wikipedia"
				/>
				{isFetching1 && showSuggestions1 ? (
					<div className="text-center text-white">Loading...</div>
				) : null}
				{showSuggestions1 && suggestions1.length > 0 && (
					<ul className="space-y-2 mb-2 bg-[#212122] px-4 py-4 rounded-md">
						{suggestions1.map((s, index) => (
							<li
								key={index}
								className="bg-[#0E1111] p-2 border-2 rounded-md border-[#7E5FFF] text-[#7E5FFF] cursor-pointer font-medium"
								onClick={() =>
									handleSuggestionClick(
										setSearchTerm1,
										setLink1,
										setShowSuggestions1,
										s
									)
								}
							>
								{s.title}
							</li>
						))}
					</ul>
				)}
			</div>

			{/* Input 2 */}
			<div className="input2">
				<div className="switch flex">
					<p className="text-white text-xl font-extrabold pt-2">To</p>
					<Image
						src={reverse}
						width={35}
						height={35}
						alt=""
						className="m-auto pb-[16px] cursor-pointer"
						onClick={swapSearchTerms}
					/>
				</div>
				<input
					type="text"
					value={searchTerm2}
					onChange={handleSearchChange2}
					onFocus={() => setShowSuggestions2(true)}
					className="w-full p-2 mb-4 text-white font-bold bg-[#212122] rounded-md focus:border-[#7E5FFF] focus:border-2 placeholder-[#7E5FFF] focus:outline-none"
					placeholder="Search Wikipedia"
				/>
				{isFetching2 && showSuggestions2 ? (
					<div className="text-center text-white">Loading...</div>
				) : null}
				{showSuggestions2 && suggestions2.length > 0 && (
					<ul className="space-y-2 mb-2 bg-[#212122] px-4 py-4 rounded-md">
						{suggestions2.map((s, index) => (
							<li
								key={index}
								className="bg-[#0E1111] p-2 border-[2px] rounded-md border-[#7E5FFF] text-[#7E5FFF] cursor-pointer font-medium"
								onClick={() =>
									handleSuggestionClick(
										setSearchTerm2,
										setLink2,
										setShowSuggestions2,
										s
									)
								}
							>
								{s.title}
							</li>
						))}
					</ul>
				)}
			</div>
			<div className="choice flex gap-10">
				<h1 className="text-[#7E5FFF] text-xl mt-5 font-black">
					Search Method:
				</h1>
				<div className="BFSDFS flex mt-2 bg-[#212122] border-[6px] border-[#212122] rounded-lg">
					<div
						className={`p-2 px-10 ${
							searchMethod === "BFS" ? "bg-[#7E5FFF]" : "bg-[#212122]"
						} text-white font-bold rounded-lg cursor-pointer text-xl`}
						onClick={toggleSearchMethod}
					>
						BFS
					</div>
					<div
						className={`p-2 px-10 ${
							searchMethod === "IDS" ? "bg-[#7E5FFF]" : "bg-[#212122]"
						} text-white font-bold rounded-md cursor-pointer ml-2 text-xl`}
						onClick={toggleSearchMethod}
					>
						IDS
					</div>
				</div>
			</div>
			<div className="Pathing flex">
				<h1 className="text-[#7E5FFF] text-xl mt-5 font-black mr-10">Path:</h1>
				<div
					className={`text-white text-xl mt-5 font-black cursor-pointer `}
					onClick={() => setPathing("SinglePath")}
					style={{ color: pathing === "SinglePath" ? "#7E5FFF" : "white" }}
				>
					Single-Path
				</div>
				<div
					className={`text-xl mt-5 font-black ml-10 cursor-pointer`}
					style={{ color: pathing === "MultiplePath" ? "#7E5FFF" : "white" }}
					onClick={() => setPathing("MultiplePath")}
				>
					Multiple-Path
				</div>
			</div>
			<div className="SetButton relative " >
				<button className="text-white bg-[#7E5FFF] py-4 pr-16 pl-8 rounded-[100px] font-black text-xl mt-10" onClick={handleSubmit}>
					Find Path
				</button>
				<Image
					src={search}
					width={35}
					height={35}
					alt="Globe"
					className="absolute top-[52px] left-[135px] transform -scale-x-100 cursor-pointer"
					onClick={handleSubmit}
				/>
			</div>
		</div>
	);
};

export default SearchInput;
