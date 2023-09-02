"use client";
import Image from "react-bootstrap/Image";
import boostchicken from "../boostchicken.png";

interface ImgProps {
  title: string;
}

const Img = (props: ImgProps) => {
  return (
    <picture>
      <Image className="logo" src={boostchicken} title={props.title} fluid />
    </picture>
  );
};
export default Img;
