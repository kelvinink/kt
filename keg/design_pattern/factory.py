#!/usr/bin/env python
# -*- coding: utf-8 -*-

class ChineseLocalizer:
    def localize(self, msg):
        return print('Chinese Localizer: ', msg)


class EnglishLocalizer:
    def localize(self, msg):
        return print('English Localizer: ', msg)


def get_localizer(language="English"):
    localizers = {
        "English": EnglishLocalizer,
        "Chinese": ChineseLocalizer,
    }
    return localizers[language]()


def main():
    en = get_localizer(language="English")
    zh = get_localizer(language="Chinese")

    en.localize('Message')
    zh.localize('Message')

if __name__ == "__main__":
    main()
