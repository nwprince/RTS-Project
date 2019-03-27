import { Component, h } from "preact";
import { tmdbResp } from "../../models/interfaces";
import Actionbar from "../actionbar";
import GenreBar from "../genrebar";
import Poster from "../poster";
import Synopsis from "../synopsis";
import Title from "../title";
import * as style from "./style.css";

interface IInfoCardProps {
  media: tmdbResp;
}

export default class InfoCard extends Component<IInfoCardProps, {}> {
  public render() {
    const genres: string[] = [];

    this.props.media.genres.forEach(element => {
      genres.push(element.name);
    });

    return (
      <div class={style.infoCardContainer}>
        <div class="display: flex">
          <Poster src={this.props.media.poster_path} />
          <div class={style.infoCardInfo}>
            <Title title={this.props.media.title} />
            <GenreBar genres={genres.slice(0, 3)} />
            <div class={style.desktop}>
              <Synopsis synopsis={this.props.media.overview} />
              <Actionbar />
            </div>
          </div>
        </div>
        <div class={style.mobile}>
          <Synopsis synopsis={this.props.media.overview} />
          <Actionbar />
        </div>
      </div>
    );
  }
}
