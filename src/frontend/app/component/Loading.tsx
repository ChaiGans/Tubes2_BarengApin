// LoadingScreen.tsx
import Image from 'next/image';
import React from 'react';
import loadingGif from "../../public/Nezuko.gif"; 

const LoadingScreen: React.FC = () => {
    return (
        <div className="w-full h-screen flex justify-center items-center bg-[#0E1111]">
            <div className="flex items-center bg-[#212122] p-5 rounded-lg shadow-lg w-3/4 max-w-4xl mx-auto border-8 border-[#7e5fff]">
                <Image 
                    src={loadingGif} 
                    alt="Loading..." 
                    width={400} 
                    height={400} 
                    className='opacity-80 rounded-2xl'
                />
                <div className="mx-auto text-white text-2xl text-[#7e5fff] font-semibold" style={{ color:"#7E5FFF"}}>
                    Loading, please wait...
                </div>
            </div>
        </div>
    );
};

export default LoadingScreen;
