import { Component, h } from "preact";
import Icon from "../icon";
import * as style from "./style.css";

export default class Actionbar extends Component {
  public render() {
    return (
      <div class={style.actionbar}>
        {Icon("info")}
        {Icon("play")}
        {Icon("share")}
      </div>
    );
  }
}
