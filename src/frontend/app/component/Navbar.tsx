import React, { RefObject } from "react";
import Link from "next/link";

export default function Navbar() {
	return (
		<div className="container flex justify-center pt-10 w-full max-w-full">
			<ul className="flex gap-10 text-white bg-[#212122] px-8 py-3 rounded-3xl font-Roboto font-black">
				<li className="cursor-pointer hover:text-[#7E5FFF]">Home</li>
				<li className="cursor-pointer hover:text-[#7E5FFF]">Docs</li>
				<li>
					<Link
						href="#services"
						className="cursor-pointer hover:text-[#7E5FFF]"
					>
						Services
					</Link>
				</li>
				<li>
					<Link
						href="#about-us"
						className="cursor-pointer hover:text-[#7E5FFF]"
					>
						About us
					</Link>
				</li>
				<li className="cursor-pointer hover:text-[#7E5FFF]">Repository</li>
			</ul>
		</div>
	);
}
