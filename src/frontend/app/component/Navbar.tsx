import React, { RefObject } from 'react';

interface NavbarProps {
    homeRef: RefObject<HTMLDivElement>;
    docsRef: RefObject<HTMLDivElement>;
    servicesRef: RefObject<HTMLDivElement>;
    aboutUsRef: RefObject<HTMLDivElement>;
    repositoryRef: RefObject<HTMLDivElement>;
}

export default function Navbar({ homeRef, docsRef, servicesRef, aboutUsRef, repositoryRef }: NavbarProps) {
    const scrollToRef = (ref: RefObject<HTMLDivElement>) => {
        if(ref.current) {
            window.scrollTo({
                top: ref.current.offsetTop,
                behavior: 'smooth'
            });
        }
    };

    return (
        <div className="container flex justify-center pt-10">
            <ul className="flex gap-10 text-white bg-[#212122] px-8 py-3 rounded-3xl font-Roboto font-black">
                <li onClick={() => scrollToRef(homeRef)} className='cursor-pointer'>Home</li>
                <li onClick={() => scrollToRef(docsRef)} className='cursor-pointer'>Docs</li>
                <li onClick={() => scrollToRef(servicesRef)} className='cursor-pointer'>Services</li>
                <li onClick={() => scrollToRef(aboutUsRef)} className='cursor-pointer'>About Us</li>
                <li onClick={() => scrollToRef(repositoryRef)} className='cursor-pointer'>Repository</li>
            </ul>
        </div>
    );
}
