import { Component, h } from "preact";
import * as style from "./style.css";

interface BackdropI {
    src: string
}

export default class Backdrop extends Component<BackdropI, {}> {
    public render() {
        return (
            <div>
                <img class={style.backdrop} width={window.innerWidth} height={window.innerHeight} src={this.props.src} />
            </div>
        );
    }
}
