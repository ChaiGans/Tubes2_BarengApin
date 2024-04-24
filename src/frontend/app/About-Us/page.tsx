/* eslint-disable react/display-name */
import Image from "next/image";
import son from "../../public/Son.jpg";
import instagram from "../../public/instagram.png";
import github from "../../public/github.png";
import React from "react";

const AboutUs = () => {
	return (
		<div>
			<div id="about-us" className="mt-20">
				<div className="Card-Page w-[100%] pt-[70px]">
					<h1 className="text-[#7E5FFF] text-5xl text-center font-black my-[20px]">
						Meet Our Developers
					</h1>
					<div className="flex justify-evenly ">
						<div className="flex justify-evenly flex-wrap gap-20">
							<div className="Card1 group relative items-center justify-center overflow-hidden cursor-pointer rounded-[20px] mx-[50px] my-[30px] border-[#7E5FFF] border-[4px]">
								<div className="h-[440px] w-64 relative ">
									<Image
										src={son}
										alt="Globe"
										className="object-cover group-hover:rotate-3 group-hover:scale-125 transition-transform h-full w-full grayscale"
									/>
								</div>
								<div className="absolute inset-0 bg-gradient-to-b from-transparent to-black/50 via-transparent to-black group-hover:from-black/70 group-hover: via-black/60 group-hover:to-black/70"></div>
								<div className="absolute inset-[-30px] flex flex-col items-center justify-center px-9 text-center translate-y-[60%] group-hover:translate-y-0 transition-all font-Merriweather">
									<h1 className="text-2xl font-bold text-[#7E5FFF]">
										Wilson Yusda
									</h1>
									<p className="text-lg text-[#7E5FFF] mb-3">13522019</p>
									<div className="Social-Media flex justify center">
										<a href="https://github.com/Razark-Y" target="_blank">
											<Image
												src={github}
												alt=""
												className="w-[40px] h-[40px] mx-10"
											/>
										</a>
										<a
											href="https://instagram.com/raflyhangga?igshid=OGQ5ZDc2ODk2ZA=="
											target="_blank"
										>
											<Image
												src={instagram}
												alt=""
												className="w-[40px] h-[40px] mr-10"
											/>
										</a>
									</div>
								</div>
							</div>
							<div className="Card2 rounded-[20px] group relative items-center justify-center overflow-hidden cursor-pointer mx-[50px] my-[30px] border-[#7E5FFF] border-[4px]">
								<div className="h-[440px] w-64 relative ">
									<Image
										src={son}
										alt="Globe"
										className="object-cover group-hover:rotate-3 group-hover:scale-125 transition-transform h-full w-full grayscale"
									/>
								</div>
								<div className="absolute inset-0 bg-gradient-to-b from-transparent to-black/50 via-transparent to-black group-hover:from-black/70 group-hover: via-black/60 group-hover:to-black/70"></div>
								<div className="absolute inset-[-20px] flex flex-col items-center justify-center px-9 text-center translate-y-[60%] group-hover:translate-y-0 transition-all font-Merriweather ">
									<h1 className="text-2xl font-bold text-[#7E5FFF]">
										Elbert Chailes
									</h1>
									<p className="text-lg text-[#7E5FFF] mb-3">13522045</p>
									<div className="Social-Media flex justify center">
										<a href="https://github.com/ChaiGans" target="_blank">
											<Image
												src={github}
												alt="Globe"
												className="w-[40px] h-[40px] mx-10"
											/>
										</a>
										<a
											href="https://instagram.com/wilson_yusda?igshid=OGQ5ZDc2ODk2ZA=="
											target="_blank"
										>
											<Image
												src={instagram}
												alt=""
												className="w-[40px] h-[40px] mr-10"
											/>
										</a>
									</div>
								</div>
							</div>
							<div className="Card3 rounded-[20px] group relative items-center justify-center overflow-hidden cursor-pointer  mx-[50px] my-[30px] border-[#7E5FFF] border-[4px]">
								<div className="h-[440px] w-64 relative ">
									<Image
										src={son}
										alt="Globe"
										className="object-cover group-hover:rotate-3 group-hover:scale-125 transition-transform h-full w-full grayscale"
									/>
								</div>
								<div className="absolute inset-0 bg-gradient-to-b from-transparent to-black/50 via-transparent to-black group-hover:from-black/70 group-hover: via-black/60 group-hover:to-black/70"></div>
								<div className="absolute inset-[-30px] flex flex-col items-center justify-center px-9 text-center translate-y-[60%] group-hover:translate-y-0 transition-all font-Merriweather">
									<h1 className="text-2xl font-bold text-[#7E5FFF]">Albert</h1>
									<p className="text-lg text-[#7E5FFF] mb-3">13522081</p>
									<div className="Social-Media flex justify center">
										<a href="https://github.com/AlbertChoe" target="_blank">
											<Image
												src={github}
												alt=""
												className="w-[40px] h-[40px] mx-10"
											/>
										</a>
										<a
											href="https://instagram.com/abdulrafirh?igshid=OGQ5ZDc2ODk2ZA=="
											target="_blank"
										>
											<Image
												src={instagram}
												alt=""
												className="w-[40px] h-[40px] mr-10"
											/>
										</a>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
};
export default AboutUs;
