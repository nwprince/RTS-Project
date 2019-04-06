import { Component, h } from "preact";
import * as style from "./style.css";

interface ISynopsisProps {
  synopsis: string;
}

export default class Synopsis extends Component<ISynopsisProps, {}> {
  public render() {
    return <div class={style.synopsis}>{this.props.synopsis}</div>;
  }
}
