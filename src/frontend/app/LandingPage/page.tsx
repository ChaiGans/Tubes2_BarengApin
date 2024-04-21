import Image from "next/image";
import globe from "../../public/Bumi.png";
import React, { forwardRef, RefObject } from 'react';
const LandingPage = forwardRef<HTMLDivElement>((props, ref) => {
    return (
        <div ref={ref} className="">
            <div className="flex items-center mt-20 pl-20">
                <div className="flex-1 mr-10">
                    <h1 className="text-5xl text-white font-black leading-[60px]">
                        Optimal Hyperlink Search: <span className="text-[#7E5FFF]">BFS</span> and <span className="text-[#7E5FFF]">IDS</span> Techniques
                    </h1>
                    <p className="text-xl text-white mt-5 leading-8">
                        Developed by a trio of university scholars, this tool employs advanced search algorithms, including Breadth-First Search (BFS) and Iterative Deepening Search (IDS), to efficiently determine the shortest hyperlink.
                    </p>
                    <button className="mt-10 text-white bg-[#7E5FFF] py-5 px-10 rounded-[100px] font-black text-2xl">
                        Get Started
                    </button>
                </div>
                <div className="flex-none mr-28">
                    <Image
                        src={globe}
                        width={400}
                        height={400}
                        alt="Globe"
                        className="spin"
                    />
                </div>
            </div>
        </div>
    );
});

export default LandingPage;
