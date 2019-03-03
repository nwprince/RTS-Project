import { Component, h } from "preact";
import * as style from "./style.css";

export default class Poster extends Component {
    public render() {
        return (
            <div>
                <img class={style.poster} src="https://image.tmdb.org/t/p/original/laMM4lpQSh5z6KIBPwWogkjzBVQ.jpg" />
            </div>
        );
    }
}