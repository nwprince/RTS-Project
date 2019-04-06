import { Component, h } from "preact";
import Backdrop from "../../components/backdrop";
import InfoCard from "../../components/infocard";
import { tmdbResp } from "../../models/interfaces";
import * as style from "./style.css";

interface IHomeProps {
  media: tmdbResp;
}

export default class Home extends Component<IHomeProps, {}> {
  public render() {
    const source: string =
      "https://image.tmdb.org/t/p/original/xFyEYxlSZ0F1lMHibVRDScY3V0P.jpg";

    return (
      <div class={style.home}>
        <Backdrop
          src={
            "https://image.tmdb.org/t/p/original/" +
            this.props.media.backdrop_path
          }
        />
        <InfoCard media={this.props.media} />
      </div>
    );
  }
}
