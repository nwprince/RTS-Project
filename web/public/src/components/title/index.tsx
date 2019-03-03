import { Component, h } from 'preact'
import Constants from '../../constants/contants';
import * as style from './style.css'

interface TitleI {
    title: string
}

export default class Title extends Component<TitleI, {}> {
    public render() {
        return (
            <div class={style.title}>{this.props.title}</div>
        );
    }
}