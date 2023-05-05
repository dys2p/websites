# On the security of the Linux disk encryption LUKS

_2023-05-03_

In the past few days, there have been uncertainties and concerns about the LUKS ("Linux Unified Key Setup") disk encryption, which is widely used on Linux. We publish our assessment of this here.

## Table of contents

* [Background](#background)
* [Discussion on Reddit](#discussion-reddit)
* [Key derivation functions](#kdf)
* [Attacks on encrypted data (media)](#attacks)
* [Recommendations](#recommendations)
  * [Upgrading encrypted drives](#upgrade)
  * [Key derivation functions of Linux distributions](#history)
  * [Reinstall instead of upgrade](#reinstall)
  * [Precautionary measures against access to unlocked devices](#unlocked-devices)
  * [Precautionary measures against attacks on passwords and passphrases](#passphrases)
* [Conclusion](#conclusion)

<h2 id="background">Background</h2>

On April 17, 2023 Matthew "mjg59" Garrett published an appeal to change the key derivation function (KDF) of LUKS-encrypted volumes: [PSA: upgrade your LUKS key derivation function](https://mjg59.dreamwidth.org/66429.html). He refers to the article "[Une lettre d’Ivan, enfermé à la prison de Villepinte : perquisitions et disques durs déchiffrés](https://nantes.indymedia.org/posts/87395/une-lettre-divan-enferme-a-la-prison-de-villepinte-perquisitions-et-disques-durs-dechiffres/)" (crosspost from [paris-luttes.info](https://paris-luttes.info/une-lettre-d-ivan-enferme-a-la-16935)).

According to the source, two encrypted computers belonging to an accused person had been decrypted – a Windows computer encrypted with BitLocker and a LUKS-encrypted computer running Ubuntu 18. (It is not clear from the text whether it is Ubuntu 18.04 or 18.10, but this is not relevant).

The Windows computer had been booted from a bootable USB stick, then a copy of the (encrypted) hard disk had been made with the software "AccessData FTK imager 3.3.05" (in the meantime "FTK® Imager by Exterro", which bought AccessData in 2020). Making a copy in order to work with it is normal in IT forensics and says nothing about the security of encryption methods. We do not want to deal with BitLocker any further at this point. However, we would like to refer to pitfalls such as vulnerabilities in implementations of the Trusted Platform Module (TPM), which enable [TPM-Sniffing-Angriffe](https://pulsesecurity.co.nz/articles/TPM-sniffing), for example, or headlines such as [Windows 10: Microsoft kriegt Bitlocker-Nachschlüssel frei Cloud (2015)](https://www.heise.de/news/Windows-10-Microsoft-kriegt-Bitlocker-Nachschluessel-frei-Cloud-3056977.html) ("Microsoft gets Bitlocker recovery key via the cloud").

The second computer, with Ubuntu 18, had been encrypted with LUKS1 using a password of more than twenty characters, consisting of letters, numbers and special characters. Garrett believes it is possible that there is no "opsec failure" and that the reason for the successful access to the Ubuntu disk is therefore the disk encryption LUKS1 and its key derivation function PBKDF2. We will go into this further below.

There could also be other causes for the decryption. For example, it is not clear from the text in which state the Ubuntu computer was seized. A seizure in a powered-on state with the screen locked or in suspend-to-RAM ("standby") allows the keys to be read from memory using a [cold boot attack](https://en.wikipedia.org/wiki/Cold_boot_attack). Since the decryption probably came as a surprise to the affected person, we would at least rule out that the password was handwritten somewhere – one of many possibilities that we will name later.

<h2 id="discussion-reddit">Discussion on Reddit</h2>

[A brief discussion](https://www.reddit.com/r/linux/comments/12q51ce/comment/jgpvsqc/) arised on Reddit between LUKS developer Clemens Fruhwirth and Matthew "mjg59" Garrett (in the Web Archive: [Fruhwirth](https://web.archive.org/web/20230420155735/https://www.reddit.com/r/linux/comments/12q51ce/comment/jgpvsqc/), [Garrett](https://web.archive.org/web/20230418093714/https://www.reddit.com/r/linux/comments/12q51ce/comment/jgq41pu/), [Fruhwirth](https://web.archive.org/web/20230418204541/https://www.reddit.com/r/linux/comments/12q51ce/comment/jgskur0/)).

Clemens Fruhwirth does not suspect LUKS1 as the cause, but other reasons. Even with LUKS1 with PBKDF2 and a password consisting of 13 random upper and lower case letters and digits, an attack with the current computing power of the entire Bitcoin network (assuming that it remains at this level and there is no further technical progress) would take 77 years. Assuming that the available computing power doubles every 24 months, it would be 10-12 years. Garrett argues that passwords are rarely truly random and thus are weaker than theoretically possible. He also notes that the computing power of graphics cards and of customized hardware (ASICs) is increasing faster than predicted.

It should be noted that Fruhwirth assumes a PBKDF2 difficulty level of 3,000,000 iterations in his calculation example. For drives that were encrypted many years ago, the difficulty level may well be lower by a factor of 20 to 30. More on this in the next section.

<h2 id="kdf">Key derivation function</h2>

Symmetric encryption methods are used to encrypt data media, for example AES ("Advanced Encryption Standard"). Symmetric encryption methods use the same cryptographic key for encryption and decryption. This key must have a certain length, for example 256 bits.

In practice, a passphrase is not used directly as a cryptographic key. Instead, a key derivation function is chosen that generates a cryptographic key from the passphrase. This has several advantages:

* The passphrase does not have to be 256 bits long, but can have a different length.
* The passphrase can be changed without having to re-encrypt the whole disk.
* We can specify the amount of time needed to generate the cryptographic key.

The last point is relevant here. To try a passphrase, the key derivation function must be executed. Accordingly, the more complex the function used, the more time-consuming it is to try a large number of possible passphrases. Modern graphics cards can execute thousands of computing operations in parallel. For this reason, key derivation functions such as **Argon2id** are now preferred, which are resistant to attacks with special hardware (graphics cards or ASICs) in that they not only require a lot of computing time, but also a lot of memory, and cannot be executed well in parallel.

When creating an encrypted volume, the difficulty level is usually chosen so that we have to wait for about one second. A volume that was encrypted many years ago on slow hardware will therefore use an older key derivation function and a lower difficulty level than a volume we encrypt today on a modern computer. Therefore, it may be wise to occasionally increase the difficulty level of the function or switch to a more modern function. This is exactly what Matthew Garrett recommends in his appeal.

The older LUKS1 only supports the key derivation function [PBKDF2](https://web.archive.org/web/20170411220929/https://www.emc.com/collateral/white-papers/h11302-pkcs5v2-1-password-based-cryptography-standard-wp.pdf), which has an adjustable computing time but requires almost no RAM. PBKDF2 is simply an adjustable number of iterations of the SHA256 or SHA512 hash algorithm. The newer LUKS2 uses [Argon2](https://en.wikipedia.org/wiki/Argon2) by default, with adjustable computation time and memory requirements. Both Garrett and, for example, the [OWASP](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html) and [KeePassXC](https://keepassxc.org/docs/KeePassXC_UserGuide) recommend the Argon2id variant. The forensics company elcomsoft [also notes](https://blog.elcomsoft.com/2022/08/probing-linux-disk-encryption-luks2-argon-2-and-gpu-acceleration/) that LUKS2 makes brute-force attacks with graphics cards considerably more difficult.

Even with LUKS1 and its key derivation function PBKDF2, there is currently no need to worry if the passphrase used is sufficiently complex.

<h2 id="attacks">Attacks on encrypted data (media)</h2>

An impression of the methods and ways to obtain or circumvent passwords and passphrases is provided, for example, in the chapter "Analyzing Filesystem Encryption" of the book "Practical Linux Forensics: A Guide for Digital Investigators":

> Decrypting protected data requires a password/passphrase or a copy of the cryptographic key (a string or key file). The forensic challenge is to find the decryption key. Some known methods for password/key recovery include:
>
> - Bruteforce with dictionarybased attacks to find simple passwords
> - Bruteforce with GPU clusters for fast exhaustive password search
> - Cryptanalysis (mathematical weakness, reduce keyspace)
> - Finding passwords saved, written, or transferred previously
> - Password reuse across multiple accounts or devices
> - Legal requirement to produce passwords in court
> - Cooperative system owner or accomplice with the password
> - Key backup/escrow in enterprise environments
> - Device exploit, vulnerability, or backdoor
> - Keyloggers or keyboard visibility (video cameras or telescope)
> - Rainbow tables: Precomputed table of cryptographic hashes
> - Extract keys from memory: PCIbus DMA attacks, hibernation
> - Maninthemiddle attacks on network traffic
> - Social engineering
> - Forced biometrics
> - Torture, blackmail, or coercion

What do others say on this topic? Mikko Hyppönen wrote in "If It's Smart, It's Vulnerable", in the chapter "Cracking Passwords":

> Authorities have pretty impressive decryption systems for this purpose. An office building in The Hague, for example, has decryption hardware the size of a supercomputer, which needs its own power station. Hardware like this allows millions of password options to be tested per second. Nevertheless, a single encrypted file can take months to open.

> Automated decryption systems use a clever tactic for speeding up this operation. If a password-protected file is found on a suspect’s hard drive, all files on the drive are indexed, and all individual words are collected from each file for testing as passwords. If none work, all discovered words are tested in reverse, and if this does not work, the drive is scanned for any unused areas and deleted files, and words inside them are tried. On surprisingly many occasions, this will decrypt the files.

However, another method offers even better chances of success:

> The best tactic is to distract a criminal, preventing them from destroying or locking their devices upon arrest. In an arrest coordinated by EUROPOL, the suspect was sitting in a café with a laptop when a female undercover officer sat at the same table, spilling her coffee a few moments later. The suspect stood up, at which point a male officer waiting behind stepped forward and whisked away the unlocked computer for forensic analysis.

The laptop of Ross Ulbricht, the operator of the darknet platform "Silk Road," [was also seized in an unlocked state](https://en.wikipedia.org/wiki/Ross_Ulbricht#Arrest). In another case, the operator of the neo-Nazi forum "Thiazi" was persuaded by a phone call to boot up his computer and unlock it shortly before he was raided.

Despite all that, the conclusion of the [OWASP Password Storage Cheat Sheets](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html) (which is aimed at online services, but can be applied to the encryption of stored data) holds true:

> Strong passwords stored with modern hashing algorithms and using hashing best practices should be effectively impossible for an attacker to crack.

<h2 id="recommendations">Recommendations</h2>

For encrypted disks in your possession, we share mjg59's recommendation to upgrade to LUKS2 and change the key derivation function to _Argon2id_:

> What makes this worse is that distributions generally don't update this (the key derivation function) in any way. If you installed your system and it gave you pbkdf2 as your KDF, you're probably still using pbkdf2 even if you've upgraded to a system that would use argon2id on a fresh install. Thankfully, this can all be fixed-up in place. But note that if anything goes wrong here you could lose access to all your encrypted data, so **before doing anything make sure it's all backed up (and figure out how to keep said backup secure so you don't just have your data seized that way).**

For data media that are out of your reach, we've put together an overview of common Linux distributions and the key derivation features they use.

<h3 id="upgrade">Upgrading encrypted drives</h3>

mjg59 wrote:

> First, make sure you're running as up-to-date a version of your distribution as possible. Having tools that support the LUKS2 format doesn't mean that your distribution has all of that integrated, and old distribution versions may allow you to update your LUKS setup without actually supporting booting from it. Also, if you're using an encrypted /boot, stop now - very recent versions of grub2 support LUKS2, but they don't support argon2id, and this will render your system unbootable.
>
> Next, figure out which device under /dev corresponds to your encrypted partition. Run
>
> `lsblk`
>
> and look for entries that have a type of "crypt". The device above that in the tree is the actual encrypted device. Record that name, and run
>
> `sudo cryptsetup luksHeaderBackup /dev/whatever --header-backup-file /tmp/luksheader`
>
> and copy that to a USB stick or something. If something goes wrong here you'll be able to boot a live image and run
>
> `sudo cryptsetup luksHeaderRestore /dev/whatever --header-backup-file luksheader`
>
> to restore it.
>
> (Edit to add: Once everything is working, delete this backup! It contains the old weak key, and someone with it can potentially use that to brute force your disk encryption key using the old KDF even if you've updated the on-disk KDF.)
>
> Next, run
>
> `sudo cryptsetup luksDump /dev/whatever`
>
> and look for the `Version:` line. If it's version 1, you need to update the header to LUKS2. Run
>
> `sudo cryptsetup convert /dev/whatever --type luks2`
>
> and follow the prompts. Make sure your system still boots, and if not go back and restore the backup of your header. Assuming everything is ok at this point, run
>
> `sudo cryptsetup luksDump /dev/whatever`
>
> again and look for the `PBKDF:` line in each keyslot (pay attention only to the keyslots, ignore any references to pbkdf2 that come after the `Digests:` line). If the PBKDF is either "pbkdf2" or "argon2i" you should convert to argon2id. Run the following:
>
> `sudo cryptsetup luksConvertKey /dev/whatever --pbkdf argon2id`
>
> and follow the prompts. If you have multiple passwords associated with your drive you'll have multiple keyslots, and you'll need to repeat this for each password.

Upgrading [persistent storage in Tails](https://tails.boum.org/doc/persistent_storage/) requires administrator privileges, which you obtain by assigning an [administration password](https://tails.boum.org/doc/first_steps/welcome_screen/administration_password/) at the login screen. Also, when upgrading from LUKS1 to LUKS2, persistent storage must not be mounted.

<h3 id="history">Key derivation functions of Linux distributions</h3>

The `cryptsetup` system tool introduced support for the LUKS2 format in [version 2.0.0](https://gitlab.com/cryptsetup/cryptsetup/-/blob/main/docs/v2.0.0-ReleaseNotes) (released 2017-12-10), using Argon2i with 128 MB memory and 800 ms computation time (on the computer used) by default. In [version 2.0.1](https://gitlab.com/cryptsetup/cryptsetup/-/blob/main/docs/v2.0.1-ReleaseNotes) (released 2018-01-21) the memory requirement was increased to 1 GiB. However, the LUKS2 format itself has only been default since [version 2.1.0](https://gitlab.com/cryptsetup/cryptsetup/-/blob/main/docs/v2.1.0-ReleaseNotes) (2019-02-08). Since [version 2.4.0](https://gitlab.com/cryptsetup/cryptsetup/-/blob/main/docs/v2.4.0-ReleaseNotes) (released 2021-08-18) cryptsetup uses Argon2id by default. However, Linux distributions often do not include the latest software versions and do not necessarily use the default options.

The relevant factor is the time of encryption, or the version of the operating system used for encryption. For example, if you encrypted a disk as part of the installation of Ubuntu 18.04 and later upgraded your operating system to Ubuntu 22.04, it will probably be encrypted with the older LUKS1. The following table shows which encryption methods are used by the **installation programs** or installation guides of the Linux distributions we analysed.

<table class="table">
  <thead>
    <tr>
      <th>Linux distribution</th>
      <th>Version of the distribution</th>
      <th>Release Date</th>
      <th>Method</th>
      <th>Discussion</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Arch Linux</td>
      <td>
        Wiki since <a href="https://wiki.archlinux.org/index.php?title=Dm-crypt/Device_encryption&amp;oldid=506702">2018-01-08</a>, default since 2019.03.01
      </td>
      <td>-</td>
      <td>LUKS2 with Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>
        Wiki since <a href="https://wiki.archlinux.org/index.php?title=Dm-crypt/Device_encryption&amp;oldid=692464">2021-08-22</a>, default since 2021.09.01
      </td>
      <td>-</td>
      <td>LUKS2 with Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Debian</td>
      <td>
        <a href="https://www.debian.org/releases/buster/amd64/release-notes/ch-whats-new.en.html#cryptsetup-luks2">since 10</a>
      </td>
      <td>2019-07-06</td>
      <td>LUKS2 with Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>11.6.0</td>
      <td>2022-12-17</td>
      <td>LUKS2 with Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td>Fedora Linux</td>
      <td>
        <a href="https://fedoraproject.org/wiki/Releases/30/ChangeSet#Switch_cryptsetup_default_metadata_format_to_LUKS2">since 30</a>
      </td>
      <td>2019-04-30</td>
      <td>LUKS2 with Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>38</td>
      <td>2023-04-18</td>
      <td>LUKS2 with Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Kali Linux</td>
      <td>2023.1</td>
      <td>2023-03-13</td>
      <td>LUKS2 with Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Linux Mint</td>
      <td>20</td>
      <td>2020-06-27</td>
      <td>LUKS2 with Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>21.1</td>
      <td>2022-12-20</td>
      <td>LUKS2 with Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Manjaro</td>
      <td>22.1.0 (tested with manjaro-xfce-22.1.0-230421)</td>
      <td>2023-04-21</td>
      <td>LUKS1 with PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td>openSUSE</td>
      <td>Leap 15.4</td>
      <td>2022-06-08</td>
      <td>LUKS1 with PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>Tumbleweed Snapshot 20230501</td>
      <td>2023-05-01</td>
      <td>LUKS1 with PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td>Pop!_OS</td>
      <td>22.04</td>
      <td>2022-04-25</td>
      <td>LUKS2 with Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td>Qubes OS</td>
      <td>
        <a href="https://www.qubes-os.org/doc/upgrade/4.1/">since 4.1</a>
      </td>
      <td>2022-02-04</td>
      <td>LUKS2 with Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>4.1.2</td>
      <td>2023-03-14</td>
      <td>LUKS2 with Argon2i</td>
      <td>
        <a href="https://forum.qubes-os.org/t/18071">Qubes forum</a>, <a href="https://www.reddit.com/r/Qubes/comments/12r4w1v/">Reddit</a>
      </td>
    </tr>
    <tr>
      <td>Tails</td>
      <td>5.12</td>
      <td>2023-04-19</td>
      <td>LUKS1 with PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>5.13 (planned)</td>
      <td>2023-05</td>
      <td>LUKS2</td>
      <td>
        <a href="https://gitlab.tails.boum.org/tails/tails/-/issues/19615">Issue 19615</a>, <a href="https://gitlab.tails.boum.org/tails/tails/-/issues/19633">Issue 19633</a>
      </td>
    </tr>
    <tr>
      <td>Ubuntu Linux</td>
      <td>18.04 LTS</td>
      <td>2018-04-23</td>
      <td>LUKS1 with PBKDF2</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>19.04</td>
      <td>2019-04-18</td>
      <td>LUKS2 with Argon2i</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>22.04.2 LTS</td>
      <td>2023-02-23</td>
      <td>LUKS2 with Argon2id</td>
      <td></td>
    </tr>
    <tr>
      <td></td>
      <td>23.04</td>
      <td>2023-04-20</td>
      <td>LUKS2 with Argon2id</td>
      <td></td>
    </tr>
  </tbody>
</table>

<h3 id="reinstall">Reinstall instead of upgrade</h3>

As Garrett also points out, it is possible that Linux installations will not support an upgrade to LUKS2 or switching to Argon2id, or will not be able to boot afterwards. To eliminate this risk, you can also backup your data on encrypted media and then reinstall your OS. (See the table above to see if this is worth doing. If you are using Tails, wait until the next release). If you upgrade encrypted disks which do not contain an operating system, you should still make a backup before upgrading them. You can also use this occasion to delete data that is no longer needed – because data that does not exist cannot be decrypted.

For those who do not want to rely on one method alone, we recommend: Encrypt your drive with LUKS2 first and store particularly vulnerable content in additional [VeraCrypt containers](https://veracrypt.fr/en/VeraCrypt%20Volume.html) with a different secure password and cipher. This way, your data is still protected even if vulnerabilities are discovered in one method. However, VeraCrypt currently also uses [PBKDF2 as its key derivation function](https://github.com/veracrypt/VeraCrypt/issues/127).

<h3 id="unlocked-devices">Precautionary measures against access to unlocked devices</h3>

By encrypting content that requires special protection separately and decrypting it only when necessary, such two-level encryption can also help against access to unlocked devices.

It can also be useful to switch off a computer, for example, by using the toggle switch of a power strip or by pulling the power plug and thus interrupting the power supply. In the case of laptops, the battery should be removed beforehand. Another way to quickly turn off or even wipe a computer is offered by devices such as [BusKill](https://www.buskill.in/luks-self-destruct/).

<h3 id="passphrases">Precautionary measures against attacks on passwords and passphrases</h3>

There are a number of countermeasures against the above-mentioned attacks on encrypted data (media). The algorithms and software used should be as up-to-date as possible (e.g. LUKS2 with Argon2id). In addition, there are mistakes that can be avoided, especially in the choice and handling of passwords and passphrases.

Choose a password or passphrase with high [entropy](https://en.wikipedia.org/wiki/Entropy_(information_theory) – in other words: so that enough possible passwords or passphrases come into question.

Methods such as [Diceware](https://en.wikipedia.org/wiki/Diceware) are ideal for this. The dice provide the necessary randomness, while the words are easy to remember using mnemonic devices. In [Creating Strong Passwords Using Dice](https://ssd.eff.org/module/creating-strong-passwords#creating-strong-passwords-using-dice), the Electronic Frontier Foundation (EFF) recommends using at least six randomly selected words from a list of 7776 words. The [EFF provides an English word list](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt) for this purpose. We have compiled a [German-language diceware word list "de-7776"](https://github.com/dys2p/wordlists-de), which will be updated again soon.

It is best to perform the procedure offline with dice. An example (with the EFF wordlist) for six words with five dice each:

1. 36433 maggot
2. 64355 unpiloted
3. 36541 manhole
4. 63656 unfunded
5. 23654 doorframe
6. 34466 implicate

The passphrase generated in the example is therefore: "maggot unpiloted manhole unfunded doorframe implicate".

The word lists are 7776 words long because there are `6^5 = 7776` different possible outcomes when rolling five six-sided dice. Each word has an entropy of `log2(7776) ≈ 13 bits`. Consequently, a passphrase of six words rolled is `5 × 13 ≈ 78 bits` strong. There are `2^78 ≈ 3 × 10^23` different possible passphrases from six words. A password consisting of thirteen randomly selected uppercase and lowercase letters and digits also has an entropy of 78 bits: one of 62 possible characters yields `log2(62) ≈ 6 bits` of entropy, thus thirteen random characters yield 78 bits.

Since the PBKDF2 key derivation function and the Bitcoin network use the same or a similar hash algorithm, their performances can be compared well. The Bitcoin network, which was used as a reference in the discussion on Reddit, currently calculates `3 × 10^20` hashes per second. For a PBKDF2 key derivation with 100,000 iterations, the network could thus try `3 × 10^15` passwords or passphrases per second. To try all possible passphrases from six words, it would take `(3 × 10^23) / (3 × 10^15) = 10^8` seconds, or about three years at constant computing power. On average, half of all possibilities must be tried, which would take about one and a half years. With a passphrase of five words instead of six, the average would be just under two hours. However, we assume that most attackers do not have such high computing power – certainly not for a long period of time. With Argon2 instead of PBKDF, both attacks should be several orders of magnitude more difficult.

Keep in mind that passwords like `IgtBa11peN!`, which are composed of the first letters of a sentence ("I go to Bed at 11pm every Night!") for example, are not random.

The online service 1Password has brought examples of passwords and the effort required to crack them in 2021: [How strong should your account password be? Here's what we learned](https://blog.1password.com/cracking-challenge-update/). However, the assumed costs may already be outdated. Further texts on how to choose passwords can be found at [The Diceware Passphrase Home Page](https://theworld.com/~reinhold/diceware.html). The FAQ from [cryptsetup](https://gitlab.com/cryptsetup/cryptsetup/-/wikis/requentlyAskedQuestions#5-security-aspects) also offers information and recommendations.

For the handling of passwords or passphrases we recommend:

* Do not use a password or passphrase and variations of it for anything else, but only for one device and use case.
* Do not write down a password or passphrase and variations of it anywhere.
* Do not share a password or passphrase and variations of it with anyone.
* Make sure that no one can secretly observe or film you when you enter it.
* Do not use biometric information.

You may also encounter "plausible deniability" methods in this context to disguise the presence of encrypted data. VeraCrypt offers the [Hidden Volume](https://www.veracrypt.fr/en/Hidden%20Volume.html) and [Hidden Operating System](https://www.veracrypt.fr/en/Hidden%20Operating%20System.html) functions for this purpose. Concerns and disadvantages on this are mentioned here, among other places:

* [Bruce Schneier: Defeating Encrypted and Deniable File Systems: TrueCrypt v5.1a and the Case of the Tattling OS and Applications](https://www.schneier.com/wp-content/uploads/2016/02/paper-truecrypt-dfs.pdf)
* [Cryptsetup: 5.18 What about Plausible Deniability?](https://gitlab.com/cryptsetup/cryptsetup/-/wikis/FrequentlyAskedQuestions#5-security-aspects)
* [Kedziora, Chow, Susilo: Defeating Plausible Deniability of VeraCrypt Hidden Operating Systems (2017)](https://ro.uow.edu.au/cgi/viewcontent.cgi?article=1542&context=eispapers1)

<h2 id="conclusion">Conclusion</h2>

In our estimation, based on the available information, PBKDF2 and LUKS1 cannot be held exclusively responsible for decryption. It is important to use a strong password or passphrase and to follow the described recommendations when using them. Using up-to-date software and algorithms increases security and reduces possible attack surfaces.

Upgrading the key derivation function is not a one-time task. You should check every few years whether the function you are using is still considered secure, and adjust the level of difficulty. This applies not only to LUKS, but to all password-based encryption tools such as VeraCrypt or KeePassXC.

## More discussions and texts

* [Systemli: Is Linux hard disk encryption hacked?](https://www.systemli.org/en/2023/04/30/is-linux-hard-disk-encryption-hacked/)
* [Discussion on Hacker News](https://news.ycombinator.com/item?id=35611425)
* [Discussion on lobste.rs](https://lobste.rs/s/ik7j1s/psa_upgrade_your_luks_key_derivation)
* [Discussion on lwn.net](https://lwn.net/Articles/929343/)
