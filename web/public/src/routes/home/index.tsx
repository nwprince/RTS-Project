import { Component, h } from "preact";
import Backdrop from "../../components/backdrop";
import Poster from "../../components/poster";
import Constants from "../../constants/contants"
import * as style from "./style.css";

export default class Home extends Component {
    public render() {
        let source: string = '';


        if (Constants.isMobile === true) {
            source = "https://image.tmdb.org/t/p/original/xFyEYxlSZ0F1lMHibVRDScY3V0P.jpg";
        } else {
            source = "https://image.tmdb.org/t/p/original/aUVCJ0HkcJIBrTJYPnTXta8h9Co.jpg";
        }


        return (
            <div class={style.home}>
                <Backdrop src={source} />
                <Poster />
            </div>
        )

    }
}
