'use client';
import Image from "next/image";
import boostchicken from "../boostchicken.svg";
import rs7 from "../boostchicken.png";
import { useRef, useEffect, useState } from "react";

const Logo = () => {

  
  const [title, setTitle] = useState("A wild boostchicken apears!");
  const src = useRef(boostchicken);
  useEffect(() => {
    if( Math.random() < 0.1) {
      src.current = rs7
      setTitle("A wild boostchicken in Ascari Blue!")
  }
  }, []);
  
  return (
    <>
      <Image className="logo" src={src.current} alt={title}  priority />
    </>
  );
};
export default Logo;
