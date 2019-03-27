import { Component, h } from "preact";
import * as style from "./style.css";

interface IPosterProps {
  src: string;
}

export default class Poster extends Component<IPosterProps, {}> {
  public render() {
    return (
      <img
        class={style.poster}
        src={"https://image.tmdb.org/t/p/original/" + this.props.src}
      />
    );
  }
}
