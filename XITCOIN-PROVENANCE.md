# Xitcoin dependency fork

Maintainer: xitcoin-org. Upstream: https://github.com/cosmos/go-ethereum/commit/d99d6fa2c8d98b7cd653de4a9386d2da3db8f25c.
Changed on 2026-09-13 for https://github.com/xitcoin-org/pos-chain/issues/33 and PR44.

NAT uses STUN v3.1.5 and retains UDP4. Go-generated module files remove DTLS v2. Loopback tests exercise successful and missing mapped addresses.

Module declarations retain their original import identity. Downstream replacements must pin an immutable revision and retain advisory mapping to upstream; a renamed fork is not a vulnerability clearance. The six-case independent security review remains required and expired on 2026-09-05.

Original copyright and license notices are preserved. SDK core: LICENSE (Apache-2.0); enterprise/group and enterprise/poa remain owned by Cosmos Labs US Inc. and governed by their respective LICENSE files (Source Available Evaluation License), only evaluated in qualification, not added to the Xitcoin application. Geth: COPYING.LESSER for library code and COPYING for GPL components/commands, subject to per-file notices. Dependencies retain their own licenses. This fork distributes corresponding source, not deployment binaries.

The fork commit is a qualification candidate until the complete required upstream checks and downstream review pass. No production suitability or global PASS is implied.
