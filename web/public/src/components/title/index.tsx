import { Component, h } from "preact";
import * as style from "./style.css";

interface ITitle {
  title: string;
}

export default class Title extends Component<ITitle, {}> {
  public render() {
    return <div class={style.title}>{this.props.title}</div>;
  }
}
