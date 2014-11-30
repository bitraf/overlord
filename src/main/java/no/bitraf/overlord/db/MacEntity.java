package no.bitraf.overlord.db;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "mac")
public class MacEntity {

    @Column(name = "account")
    protected int account;

    @Id
    @Column(name = "macaddr")
    protected String macAddr;

    @Column(name = "device")
    protected String device;

    public int getAccount() {
        return account;
    }

    public String getMacAddr() {
        return macAddr;
    }

    public String getDevice() {
        return device;
    }
}
