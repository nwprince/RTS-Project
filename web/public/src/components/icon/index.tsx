import { h } from "preact";
import Info from "./info";
import Play from "./play";
import Share from "./share";

const Icon = (name: string) => {
  switch (name) {
    case "info":
      return <Info />;

    case "play":
      return <Play />;

    case "share":
      return <Share />;

    default:
      break;
  }
};

export default Icon;
