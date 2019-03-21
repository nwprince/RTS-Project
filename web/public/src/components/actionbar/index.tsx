import { Component, h } from "preact";
import * as style from "./style.css";

export default class Actionbar extends Component {
  public render() {
    return (
      <div class={style.actionbar}>
        <button>Info</button>
        <button>Play</button>
        <button>Share</button>
      </div>
    );
  }
}
