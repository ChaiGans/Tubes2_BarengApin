// @ts-nocheck
"use client";
import React, { useState, useEffect, useRef } from "react";
import Navbar from "./component/Navbar";
import LandingPage from "./LandingPage/page";
import AboutUs from "./About-Us/page";
import Search from "./component/SearchInput";
import * as d3 from "d3";
interface Node {
	id: string;
	group: number;
	x?: number;
	y?: number;
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
	const [searchResult, setSearchResult] = useState<ForceGraphProps | null>(
		null
	);
	const [Duration, setDuration] = useState(null);
	const [Timevisited, setTimevisited] = useState(null);
	const homeRef = useRef(null);
	const [Path, setNumPath] = useState(null);
	const docsRef = useRef(null);
	const servicesRef = useRef(null);
	const aboutUsRef = useRef(null);
	const repositoryRef = useRef(null);
	const handleSearchResult = (result: any) => {
		console.log(result);
		const d3FormattedData = transformResultToD3Format(JSON.stringify(result));
		setNumPath(result.shortestPath.length);
		setSearchResult(d3FormattedData);
		setDuration(result.exectime);
		setTimevisited(result.numchecked);
	};
	useEffect(() => {
		if (searchResult) {
			console.log("Search result:", searchResult);
		}
	}, [searchResult]);
	function transformResultToD3Format(resultJson: string): ForceGraphProps {
		const result = JSON.parse(resultJson);
		const paths: string[][] = result.shortestPath;
		console.log(paths);
		const nodeSet = new Map<string, number>(); 
		const links: Link[] = [];
		paths.forEach((path: string[]) => {
			path.forEach((url, index) => {
				const nodeName = decodeURIComponent(
					new URL(url).pathname.split("/").pop()!
				);
				if (!nodeSet.has(nodeName) || nodeSet.get(nodeName)! > index) {
					nodeSet.set(nodeName, index);
				}
	
				if (index < path.length - 1) {
					const nextUrl = path[index + 1];
					const nextNodeName = decodeURIComponent(
						new URL(nextUrl).pathname.split("/").pop()!
					);
					links.push({
						source: nodeName,
						target: nextNodeName,
						value: 1, 
					});
				}
			});
		});
	
		const nodes: Node[] = Array.from(nodeSet).map(([nodeId, depth]) => ({
			id: nodeId,
			group: depth
		}));
	
		return { nodes, links };
	}
	
	const ForceGraph: React.FC<ForceGraphProps> = ({ nodes, links }) => {
		const svgRef = useRef<SVGSVGElement>(null);
	
		useEffect(() => {
			if (!svgRef.current) return;
			const width = 1365; 
			const height = 700;
			const svg = d3.select(svgRef.current);
			const depthColorMapping = {
				0: '#7e5fff', 
				1: '#FD1C03', 
				2: '#FF1178', 
				3: '#E6FB04', 
				4: '#099FFF', 
			};
			svg.selectAll("*").remove();
	
			svg.append("defs")
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
	
			const zoom = d3.zoom()
				.scaleExtent([0.5, 5])
				.on("zoom", (event) => {
					content.attr("transform", event.transform);
				});
	
			const content = svg.append("g").attr("class", "content");
	
			const depthIndices = {};
			nodes.forEach(node => {
				if (depthIndices[node.group] === undefined) depthIndices[node.group] = 0;
				else depthIndices[node.group]++;
				node.x = 150 + 200 * node.group; 
				node.y = 100 + 100 * depthIndices[node.group]; 
			});
	
			const simulation = d3.forceSimulation(nodes)
				.force("link", d3.forceLink(links).id((d) => d.id).distance(400))
				.force("charge", d3.forceManyBody())
				.force("center", d3.forceCenter(width / 2, height / 2));
	
			const link = content.append("g")
				.selectAll("line")
				.data(links)
				.join("line")
				.attr("stroke", "#7e5fff")
				.attr("stroke-width", (d) => Math.sqrt(d.value))
				.attr("marker-end", "url(#end)");
	
			const nodeRadius = 20;
			const node = content.append("g")
				.selectAll("circle")
				.data(nodes)
				.join("circle")
				.attr("r", nodeRadius)
				.attr("fill", d => depthColorMapping[d.group] || '#000')
				.attr("stroke", "#0E1111")
				.attr("stroke-width", 5)
				.call(drag(simulation));
	
			const labels = content.append("g")
				.selectAll("text")
				.data(nodes)
				.join("text")
				.text((d) => d.id)
				.attr("x", (d) => d.x + nodeRadius + 5)
				.attr("y", (d) => d.y + nodeRadius / 2)
				.attr("fill", "white")
				.style("font-size", "18px")
				.style("pointer-events", "none")
				.style("font-weight", "bold");
	
			const maxDepth = Math.max(...nodes.map(n => n.group));
	
			const usedDepths = Object.entries(depthColorMapping)
				.filter(([depth]) => depth <= maxDepth);
	
			const legend = svg.append("g")
				.attr("class", "legend")
				.attr("transform", `translate(30, ${height - 50})`) 
				.selectAll("g")
				.data(usedDepths)
				.enter()
				.append("g")
				.attr("transform", (d, i) => `translate(${i * 150}, 0)`); 
	
			legend.append("rect")
				.attr("width", 20)
				.attr("height", 20)
				.attr("fill", (d) => d[1]);
	
			legend.append("text")
				.attr("x", 30)
				.attr("y", 15)
				.text((d) => `Depth ${d[0]}`)
				.attr("fill", "white")
				.style("font-size", "16px")
				.style("font-weight", "bold");
	
			svg.call(zoom);
	
			simulation.on("tick", () => {
				link.attr("x1", (d) => d.source.x)
					.attr("y1", (d) => d.source.y)
					.attr("x2", (d) => d.target.x)
					.attr("y2", (d) => d.target.y);
	
				node.attr("cx", (d) => d.x).attr("cy", (d) => d.y);
	
				labels.attr("x", (d) => d.x + nodeRadius + 10)
					.attr("y", (d) => d.y + nodeRadius / 2 + 5);
			});
	
			function drag(simulation) {
				return d3.drag()
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
			<Navbar />
			<LandingPage></LandingPage>
			<AboutUs></AboutUs>
			<Search onSearchResult={handleSearchResult} id="services"></Search>
			<div className="runtime-info text-white text-xl ml-20 mr-20 mt-20 px-10 mb-10 flex justify-between">
				<p className="text-[#7e5fff] font-semibold">
					<strong>Runtime:</strong> {Duration !== null ? `${Duration} ms` : "-"}
				</p>
				<p className="text-[#7e5fff] font-semibold">
					<strong>Links Visited:</strong>{" "}
					{Timevisited !== null ? `${Timevisited} links` : "-"}
				</p>
				<p className="text-[#7e5fff] font-semibold">
					<strong>Number of Path:</strong>{" "}
					{Path !== null ? `${Path} path` : "-"}
				</p>
			</div>
			<div className="Graph bg-[#212122] w-[90%] h-[700px] m-auto rounded-xl ml-20 mr-20 mb-10">
				{searchResult && <ForceGraph {...searchResult} />}
			</div>
		</main>
	);
}
