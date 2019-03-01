import { Component, h } from "preact";
import { Link } from "preact-router/match";
import Constants from "../../constants/contants"
import * as style from "./style.css";

export default class Header extends Component {
    public render() {
        if (Constants.isMobile === true) {
            return <header class={style.mobile}>
                <h1>GoMedia</h1>
            </header>
        }
        return (
            <header class={style.desktop}>
                <h1>GoMedia</h1>
            </header>
        );
    }
}
