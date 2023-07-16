from soco import SoCo, discover

def main():
    zones = list(discover())
    print(zones)

    zones[0].play_uri(
        'http://ia801402.us.archive.org/20/items/TenD2005-07-16.flac16/TenD2005-07-16t10Wonderboy.mp3'
    )
    # zones[0].pause()

