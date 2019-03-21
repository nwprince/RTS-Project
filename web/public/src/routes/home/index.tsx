import { Component, h } from "preact";
import Backdrop from "../../components/backdrop";
import InfoCard from "../../components/infocard";
import * as style from "./style.css";

export default class Home extends Component {
  public render() {
    const source: string =
      "https://image.tmdb.org/t/p/original/xFyEYxlSZ0F1lMHibVRDScY3V0P.jpg";

    return (
      <div class={style.home}>
        <Backdrop src={source} />
        <InfoCard />
      </div>
    );
  }
}
