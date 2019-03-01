import { Component, h } from "preact";
import Constants from "../../constants/contants"
import * as style from "./style.css";


export default class Poster extends Component {

    public render() {
        if (Constants.isMobile) {
            return (
                <div>
                    <img class={style.mobile} src="https://image.tmdb.org/t/p/original/laMM4lpQSh5z6KIBPwWogkjzBVQ.jpg" />
                </div>
            );
        } 
            return (
                <div>
                    <img class={style.desktop} src="https://image.tmdb.org/t/p/original/laMM4lpQSh5z6KIBPwWogkjzBVQ.jpg" />
                </div>
            );
        


    }
}