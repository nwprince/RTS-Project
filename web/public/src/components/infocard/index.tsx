import { Component, h } from "preact";
import Actionbar from "../actionbar";
import GenreBar from "../genrebar";
import Poster from "../poster";
import Synopsis from "../synopsis";
import Title from "../title";
import * as style from "./style.css";

export default class InfoCard extends Component {
  public render() {
    return (
      <div class={style.infoCardContainer}>
        <div class="display: flex">
          <Poster />
          <div class={style.infoCardInfo}>
            <Title title="Spiderman: Into the Spiderverse" />
            <GenreBar genres={["Animation", "Action", "Comedy"]} />
            <div class={style.desktop}>
              <Synopsis synopsis="Miles Morales is juggling his life between being a high school student and being a spider-man. When Wilson Kingpin Fisk uses a super collider, others from across the Spider-Verse are transported to this dimension." />
              <Actionbar />
            </div>
          </div>
        </div>
        <div class={style.mobile}>
          <Synopsis synopsis="Miles Morales is juggling his life between being a high school student and being a spider-man. When Wilson Kingpin Fisk uses a super collider, others from across the Spider-Verse are transported to this dimension." />
          <Actionbar />
        </div>
      </div>
    );
  }
}
