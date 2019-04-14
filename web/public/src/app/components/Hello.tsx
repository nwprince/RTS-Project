import * as React from 'react';

const Ipfs = require('ipfs');
const Hls = require('hls.js');
const HlsjsIpfsLoader = require('hlsjs-ipfs-loader');
const testhash = 'QmYB65ugMfUc4axC8vNmoL6fTPio7PFJYoNJzn74uKGR44';
const ipfs = new Ipfs({ repo: 'ipfs-' + Math.random() });

export class Hello extends React.Component {
	private hls: any;
	private video: any;
	constructor(props: any, context: any) {
		super(props, context);
		Hls.DefaultConfig.loader = HlsjsIpfsLoader;
		Hls.DefaultConfig.debug = false;
		this.hls = new Hls();
	}

	componentDidMount() {
		ipfs.on('ready', () => {
			console.log('success');
			if (Hls.isSupported()) {
				const video = document.getElementById('video');
				this.hls.config.ipfs = ipfs;
				this.hls.config.ipfsHash = testhash;
				this.hls.loadSource('master.m3u8');
				this.hls.attachMedia(video);
				this.hls.on(Hls.Events.MANIFEST_PARSED, () => this.video.play());
			}
		});
	}

	componentWillUnmount() {
		if (this.hls) {
			this.hls.destroy();
		}
	}

	render() {
		return <div />;
	}
}
