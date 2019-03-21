import { Component, h } from "preact";
import * as style from "./styles.css";

interface IGenreBar {
  genres: string[];
}

export default class GenreBar extends Component<IGenreBar, {}> {
  public render() {
    const genres = this.props.genres.join(" • ");
    return <div class={style.item}>{genres}</div>;
  }
}
