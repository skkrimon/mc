#!/usr/bin/python3

import requests
import os
import configparser
from argparse import ArgumentParser
from urllib import request
from time import sleep, time
from tqdm import tqdm
from yaspin import yaspin
from pathlib import Path

RESTART_TIME = 60
BASE_URL = 'https://api.papermc.io/v2'
FILENAME = 'server.jar'
SERVER_PATH = '../../server'
BACKUP_PATH = '../../backup'
START_SERVER_CMD = 'systemctl start minecraft'
STOP_SERVER_CMD = 'systemctl stop minecraft'
CONFIG_FILE = 'update.ini'
USER = 'minecraft'
GROUP = 'minecraft'


class DownloadProgressBar(tqdm):
    def update_to(self, b=1, bsize=1, tsize=None):
        if tsize is not None:
            self.total = tsize
        self.update(b * bsize - self.n)


def create_config():
    filename = Path(CONFIG_FILE)
    filename.touch(exist_ok=True)


@yaspin(text='Fetching latest version info')
def get_latest_version():
    versions_url = f'{BASE_URL}/projects/paper'
    versions = requests.get(versions_url).json()['versions']
    return versions.pop()


@yaspin(text='Fetching latest build info')
def get_latest_build(v):
    builds_url = f'{BASE_URL}/projects/paper/versions/{v}/builds'
    builds = requests.get(builds_url).json()['builds']
    return builds.pop()


def download(url):
    print(f'Downloading {FILENAME}')

    with DownloadProgressBar(unit='B', unit_scale=True, miniters=1, desc=url.split('/')[-1]) as t:
        request.urlretrieve(url, filename=FILENAME, reporthook=t.update_to)


@yaspin(text='Shutting down server')
def stop_server():
    os.system(STOP_SERVER_CMD)
    sleep(RESTART_TIME)


@yaspin(text='Restarting server')
def start_server():
    os.system(START_SERVER_CMD)
    sleep(RESTART_TIME)


if __name__ == '__main__':
    try:
        parser = ArgumentParser()
        parser.add_argument('-v', '--version')
        parser.add_argument('-b', '--build')
        parser.add_argument('-d', '--download')

        args = parser.parse_args()

        version = args.version
        build = args.build
        download_file = args.download

        if version is None:
            version = get_latest_version()

        if build is None:
            build = get_latest_build(version)['build']

        if download_file is None:
            download_file = f'paper-{version}-{build}.jar'

        create_config()

        config = configparser.ConfigParser()
        config.read(CONFIG_FILE)

        try:
            current_version = config['Settings']['current_version']
            current_build = config['Settings']['current_build']
        except KeyError as e:
            current_version = ''
            current_build = ''

        if str(current_build) == str(build) and str(current_version) == str(version):
            print('There is no update available.')
            print(f'Current version {current_version}')
            print(f'Current build: {current_build}')
            exit()

        download_url = f'{BASE_URL}/projects/paper/versions/{version}/builds/{build}/downloads/{download_file}'
        download(download_url)

        stop_server()

        timestamp = time()

        os.system(f'mv {SERVER_PATH}/{FILENAME} {BACKUP_PATH}/{timestamp}_{FILENAME}')
        print('Backed up old server.jar version')

        os.system(f'mv {FILENAME} {SERVER_PATH}/{FILENAME}')
        print('Installed new server.jar version')

        os.system(f'chown {USER}:{GROUP} {SERVER_PATH}/{FILENAME}')

        start_server()

        print('\nSuccessfully installed latest paper server')

        config['Settings'] = {
            'current_version': version,
            'current_build': build
        }

        with open(CONFIG_FILE, 'w') as configfile:
            config.write(configfile)

        print('Updated update.ini')

    except Exception as e:
        print(f'An error occurred while fetching {e}\n')
        exit()
