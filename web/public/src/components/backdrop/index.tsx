import { Component, h } from "preact";
import Constants from "../../constants/contants";
import * as style from "./style.css";

interface BackdropI {
    src: string
}

export default class Backdrop extends Component<BackdropI, {}> {
    public render() {
        if (Constants.isMobile === true) {
            return (
                <div>
                    <img width={window.innerWidth} height={window.innerHeight} src={this.props.src} />
                    <div class={style.gradient} />
                    <div class={style.gradient} />
                </div>
            );
        }

        return (
            <img class={style.desktop} width={window.innerWidth} height={window.innerHeight} src={this.props.src} />
        );
    }
}
