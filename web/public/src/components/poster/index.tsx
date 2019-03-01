import { Component, h } from "preact";
import Constants from "../../constants/contants"
import * as style from "./style.css";

export default class Poster extends Component {
    public render() {
        let posterWidth: number = 0;
        let posterHeight: number = 0;
        let className: string;

        if (Constants.isMobile) {
            posterWidth = window.innerWidth;
            posterHeight = window.innerWidth;
            className = style.mobile;
        } else {
            posterWidth = 0.4 * window.innerWidth;
            posterHeight = window.innerHeight;
            className = style.desktop
        }

        return (
            <div>
                <img class={className} width={posterWidth} height={posterHeight} src="https://image.tmdb.org/t/p/original/laMM4lpQSh5z6KIBPwWogkjzBVQ.jpg" />
            </div>
        );
    }
}