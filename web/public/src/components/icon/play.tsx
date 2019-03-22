import { h } from "preact";

const Play = ({
  style = {},
  fill = "#fff",
  width = "5vw",
  height = "5vh",
  viewBox = "0 0 48 48"
}) => {
  return (
    <svg
      width={width}
      style={style}
      height={height}
      viewBox={viewBox}
      xmlns="http://www.w3.org/2000/svg"
      xmlnsXlink="http://www.w3.org/1999/xlink"
    >
      <path d="M0 0h48v48H0z" fill="none" />
      <path
        fill={fill}
        d="M20 33l12-9-12-9v18zm4-29C12.95 4 4 12.95 4 24s8.95 20 20 20 20-8.95 20-20S35.05 4 24 4zm0 36c-8.82 0-16-7.18-16-16S15.18 8 24 8s16 7.18 16 16-7.18 16-16 16z"
      />
    </svg>
  );
};

export default Play;
