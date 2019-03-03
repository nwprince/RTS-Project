import { Component, h } from "preact";
import GenreBar from "../genrebar";
import Title from "../title";
import * as style from './style.css';

export default class InfoCard extends Component {
    public render() {
        return (
            <div class={style.infobar}>
                <Title title="Spiderman: Into the Spiderverse" />
                <GenreBar genres={["Animation", "Action", "Comedy"]} />
            </div>
        );
    }
}