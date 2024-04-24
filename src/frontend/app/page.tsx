"use client";
import Image from "next/image";
import React, { useState, useEffect, useRef } from "react";
import Form from "./Form/page";
import Navbar from "./component/Navbar";
import LandingPage from "./LandingPage/page";
import AboutUs from "./About-Us/page";
import Search from "./component/SearchInput";
import * as d3 from "d3";
interface Node {
	id: string;
	group: number;
}

interface Link {
	source: string;
	target: string;
	value: number;
}

interface ForceGraphProps {
	nodes: Node[];
	links: Link[];
}
export default function Home() {
	const [searchResult, setSearchResult] = useState(null);
	const homeRef = useRef(null);
	const docsRef = useRef(null);
	const servicesRef = useRef(null);
	const aboutUsRef = useRef(null);
	const repositoryRef = useRef(null);
	const handleSearchResult = (result) => {
		console.log(result);
		const d3FormattedData = transformResultToD3Format(JSON.stringify(result));
		setSearchResult(d3FormattedData);
	};
	useEffect(() => {
		if (searchResult) {
			console.log("Search result:", searchResult);
		}
	}, [searchResult]);
	// const miserables_data = {
	//     nodes: [
	// 		{"id": "Valjean", "group": 1},
	// 		{"id": "Javert", "group": 1},
	// 		{"id": "Fantine", "group": 2},
	// 		{"id": "Cosette", "group": 2},
	// 		{"id": "Marius", "group": 3},
	// 		{"id": "Enjolras", "group": 3},
	// 		{"id": "Bishop", "group": 5},
	// 		{"id": "Myriel", "group": 5},
	//     ],
	//     links: [
	// 		{"source": "Valjean", "target": "Javert", "value": 5},
	// 		{"source": "Valjean", "target": "Fantine", "value": 5},
	// 		{"source": "Fantine", "target": "Cosette", "value": 5},
	// 		{"source": "Valjean", "target": "Cosette", "value": 5},
	// 		{"source": "Cosette", "target": "Marius", "value": 5},
	// 		{"source": "Valjean", "target": "Marius", "value": 5},
	// 		{"source": "Marius", "target": "Enjolras", "value": 5},
	// 		{"source": "Valjean", "target": "Bishop", "value": 5},
	// 		{"source": "Bishop", "target": "Myriel", "value": 5},
	//     ],
	// };
	function transformResultToD3Format(resultJson: string): ForceGraphProps {
		const result = JSON.parse(resultJson);
		const paths: string[][] = result.shortestPath;
		console.log(paths);  
		const nodeSet = new Set<string>();
		const links: Link[] = [];
		paths.forEach((path: string[]) => {
		  path.forEach((url, index) => {
			const nodeName = decodeURIComponent(new URL(url).pathname.split("/").pop()!);
			nodeSet.add(nodeName);
	  
			if (index < path.length - 1) {
			  const nextUrl = path[index + 1];
			  const nextNodeName = decodeURIComponent(new URL(nextUrl).pathname.split("/").pop()!);
			  links.push({
				source: nodeName,
				target: nextNodeName,
				value: 1
			  });
			}
		  });
		});
		const nodes: Node[] = Array.from(nodeSet).map(nodeId => ({ id: nodeId, group: 1 }));
		return { nodes, links };
	  }
	const ForceGraph: React.FC<ForceGraphProps> = ({ nodes, links }) => {
		const svgRef = useRef<SVGSVGElement>(null);

		useEffect(() => {
			if (!svgRef.current) return;
			const width = 1365; // Full width of the container
			const height = 700;
			const svg = d3.select(svgRef.current);
			svg.selectAll("*").remove();
			svg
				.append("defs")
				.selectAll("marker")
				.data(["end"])
				.enter()
				.append("marker")
				.attr("id", (d) => d)
				.attr("viewBox", "0 -5 10 10")
				.attr("refX", 25)
				.attr("refY", 0)
				.attr("markerWidth", 16)
				.attr("markerHeight", 16)
				.attr("orient", "auto")
				.append("path")
				.attr("fill", "#7e5fff")
				.attr("d", "M0,-5L10,0L0,5");
			const zoom = d3
				.zoom()
				.scaleExtent([0.5, 5])
				.on("zoom", (event) => {
					content.attr("transform", event.transform);
				});

			const content = svg.append("g").attr("class", "content");

			nodes.forEach((node) => {
				node.x = width / 2;
				node.y = height / 2;
			});

			const simulation = d3
				.forceSimulation(nodes)
				.force(
					"link",
					d3
						.forceLink(links)
						.id((d) => d.id)
						.distance(400)
				)
				.force("charge", d3.forceManyBody())
				.force("center", d3.forceCenter(width / 2, height / 2));

			const link = content
				.append("g")
				.selectAll("line")
				.data(links)
				.join("line")
				.attr("stroke", "#7e5fff")
				.attr("stroke-width", (d) => Math.sqrt(d.value))
				.attr("marker-end", "url(#end)");

			const nodeRadius = 20;
			const node = content
				.append("g")
				.selectAll("circle")
				.data(nodes)
				.join("circle")
				.attr("r", nodeRadius)
				.attr("fill", "0e1111")
				.attr("stroke", "#7e5fff") // Node border color
				.attr("stroke-width", 2) // Node border width
				.call(drag(simulation));

			const labels = content
				.append("g")
				.selectAll("text")
				.data(nodes)
				.join("text")
				.text((d) => d.id)
				.attr("x", (d) => d.x + nodeRadius + 5)
				.attr("y", (d) => d.y + nodeRadius / 2)
				.attr("fill", "white") // Set text color to white
				.style("font-size", "18px") // Optional: Set font size
				.style("pointer-events", "none")
				.style("font-weight", "bold");
			svg.call(zoom);

			simulation.on("tick", () => {
				link
					.attr("x1", (d) => d.source.x)
					.attr("y1", (d) => d.source.y)
					.attr("x2", (d) => d.target.x)
					.attr("y2", (d) => d.target.y);

				node.attr("cx", (d) => d.x).attr("cy", (d) => d.y);

				labels
					.attr("x", (d) => d.x + nodeRadius + 10)
					.attr("y", (d) => d.y + nodeRadius / 2 + 5);
			});

			function drag(simulation) {
				return d3
					.drag()
					.on("start", (event) => {
						if (!event.active) simulation.alphaTarget(0.3).restart();
						event.subject.fx = event.subject.x;
						event.subject.fy = event.subject.y;
					})
					.on("drag", (event) => {
						event.subject.fx = event.x;
						event.subject.fy = event.y;
					})
					.on("end", (event) => {
						if (!event.active) simulation.alphaTarget(0);
						event.subject.fx = null;
						event.subject.fy = null;
					});
			}

			return () => simulation.stop();
		}, [nodes, links]);

		return <svg ref={svgRef} width="100%" height="100%" />;
	};
	return (
		<main className=" min-h-screen bg-[#0E1111] pb-20">
			<Navbar
				homeRef={homeRef}
				docsRef={docsRef}
				servicesRef={servicesRef}
				aboutUsRef={aboutUsRef}
				repositoryRef={repositoryRef}
			/>
			<LandingPage ref={homeRef}></LandingPage>
			<AboutUs ref={aboutUsRef}></AboutUs>
			<Search onSearchResult={handleSearchResult} ref={servicesRef}></Search>
			<div className="Graph bg-[#212122] w-[90%] h-[700px] m-auto rounded-xl ml-20 mr-20 mt-20 mb-10">
				{searchResult && <ForceGraph {...searchResult} />}
			</div>
		</main>
	);
}
