import { Component, h } from "preact";
import * as style from "./style.css";

interface IBackdrop {
  src: string;
}

export default class Backdrop extends Component<IBackdrop, {}> {
  public render() {
    if (window.innerWidth <= 500) {
      return (
        <div class={style.outer}>
          <div class={style.inner}>
            <img
              // width={window.innerWidth}
              // height={window.innerHeight}
              src={this.props.src}
            />
          </div>
          <div class={style.gradient} />
        </div>
      );
    }

    return (
      <img
        class={style.desktop}
        width={window.innerWidth}
        height={window.innerHeight}
        src={this.props.src}
      />
    );
  }
}
